package main

import (
	"accounting_software/telegram"
	"accounting_software/telegram/post"
	"bytes"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const tgbotapiKey = "962901703:AAG2Yztqpf91ZISxm2oZElHrBUlk6hSi_IU"

const postURL = "http://localhost:3001/order"
const postCustomerSetURL = "http://localhost:3001/customer/add"
const postCustomerChatIDURL = "http://localhost:3001/customer/get/chat_id"
const postCustomerLastIDURL = "http://localhost:3001/customer/get/last_id"
const postRemoveOrderURL = "http://localhost:3001/order/remove"
const postGetLastOrderURL = "http://localhost:3001/order/last"

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Заказать"),
	),
)
var mainMenuSecondOrder = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Заказать"),
		tgbotapi.NewKeyboardButton("Отменить последний заказ"),
	),
)
var productMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🥖 батон"),
		tgbotapi.NewKeyboardButton("🍞 кирпич"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🥐 булка"),
		tgbotapi.NewKeyboardButton("🍞 ржаной"),
	),
)
var amountMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("3"),
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
	),
)

var orderMap map[int]*telegram.Order

//pered startom programmi
func init() {
	orderMap = make(map[int]*telegram.Order)
}

func main() {
	var (
		bot        *tgbotapi.BotAPI
		err        error
		updChannel tgbotapi.UpdatesChannel
		update     tgbotapi.Update
		updConfig  tgbotapi.UpdateConfig
		botUser    tgbotapi.User

		itemCustomer *telegram.Customer
		itemOrder    *telegram.OrderResp
	)

	client := http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout:   time.Second * 10,
			DisableKeepAlives:     false,
			DisableCompression:    false,
			IdleConnTimeout:       time.Second * 10,
			ResponseHeaderTimeout: time.Second * 10,
			ExpectContinueTimeout: time.Second * 10,
		},
		Timeout: time.Second * 10,
	}

	//inicaleziruem bota
	bot, err = tgbotapi.NewBotAPI(tgbotapiKey)
	if err != nil {
		fmt.Printf("Bot init error: %s\n", err)
		return
	}

	botUser, err = bot.GetMe()
	if err != nil {
		fmt.Printf("Bot getme error: %s\n", err)
		return
	}

	fmt.Printf("auth ok! Bot is: %s\n ", botUser.FirstName)
	//poluchili canal

	updConfig.Timeout = 60
	updConfig.Limit = 10
	updConfig.Offset = 0

	updChannel, err = bot.GetUpdatesChan(updConfig)
	if err != nil {
		fmt.Printf("Update channel error: %s\n", err)
		return
	}

	//tgbotapi.InlineKeyboardMarkup{}
	//btn := tgbotapi.KeyboardButton{Text: "Заказать товара"}
	for {
		update = <-updChannel
		if update.Message != nil {
			//esli comanda:
			params := fmt.Sprintf(
				"{\"ChatID\": %d}",
				update.Message.Chat.ID)
			buf := bytes.NewBufferString(params)

			//////////////////////////////POST GET CHATID//////////////////////////////////////////
			req, err := http.NewRequest("POST", postCustomerChatIDURL, buf)
			req.Header.Set("content-type", "application/json")
			if err != nil {
				panic(err)
			}

			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			//defer resp.Body.Close()
			itemCustomer = new(telegram.Customer)
			body, err := ioutil.ReadAll(resp.Body)
			err = json.Unmarshal(body, itemCustomer)
			if err != nil {
				fmt.Printf("Error Unmarshal ChatID")
				return
			}

			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				if cmdText == "start" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Здравствуйте! Для заказа нажмите кнопку заказать ⬇")
					msg.ReplyMarkup = mainMenu
					bot.Send(msg)
				}
			} else {
				//if chatid exist in db
				if itemCustomer.CustomerName != "" {
					//--------------------------------------------------------------------------------------------------
					//---------------------------------------------OLD USER-------------------------------------------
					//--------------------------------------------------------------------------------------------------
					//esli ne komanda:
					if update.Message.Text == mainMenuSecondOrder.Keyboard[0][0].Text {
						orderMap[update.Message.From.ID] = new(telegram.Order)
						orderMap[update.Message.From.ID].State = 0
						msgConfig := tgbotapi.NewMessage(
							update.Message.Chat.ID,
							"Выберите пожалуйста продукт который Вы желаете заказать в магазине:")
						msgConfig.ReplyMarkup = productMenu
						bot.Send(msgConfig)
					} else if update.Message.Text == mainMenuSecondOrder.Keyboard[0][1].Text {
						fmt.Printf("remove order")

						//////////////////////////SEND post Last Order////////////////////////
						params := fmt.Sprintf(
							"{\"CustomerID\": %d}",
							itemCustomer.ID)
						buf := bytes.NewBufferString(params)
						req, err := http.NewRequest("POST", postGetLastOrderURL, buf)
						req.Header.Set("content-type", "application/json")
						if err != nil {
							panic(err)
						}
						resp, err := client.Do(req)
						if err != nil {
							panic(err)
						}
						itemOrder = new(telegram.OrderResp)
						body, err := ioutil.ReadAll(resp.Body)
						err = json.Unmarshal(body, itemOrder)
						if err != nil {
							fmt.Printf("Error Unmarshal")
							return
						}
						fmt.Printf("Error Unmarshal %s\n", itemOrder.Product)

						if itemOrder.Product != "" {
							//////////////////////////SEND post REMOVE ORDER////////////////////////
							params1 := fmt.Sprintf(
								"{\"ID\": %d}",
								itemCustomer.ID)
							buf1 := bytes.NewBufferString(params1)
							req1, err := http.NewRequest("POST", postRemoveOrderURL, buf1)
							req1.Header.Set("content-type", "application/json")
							if err != nil {
								panic(err)
							}
							resp1, err := client.Do(req1)
							if err != nil {
								panic(err)
							}
							fmt.Printf("%v\n", resp1)

							msgConfig := tgbotapi.NewMessage(
								update.Message.Chat.ID,
								"Ваш заказ \""+itemOrder.Product+"\" - "+strconv.Itoa(itemOrder.Amount)+"шт. был успешно отменен!",
							)
							msgConfig.ReplyMarkup = mainMenuSecondOrder
							bot.Send(msgConfig)
						} else {
							msgConfig := tgbotapi.NewMessage(
								update.Message.Chat.ID,
								"Больше нет заказов!",
							)
							msgConfig.ReplyMarkup = mainMenu
							bot.Send(msgConfig)
						}
					} else {
						ord, ok := orderMap[update.Message.From.ID]
						if ok {
							if ord.State == 0 {
								ord.Product = update.Message.Text
								ord.ProductID = productId(ord.Product)
								msgConfig := tgbotapi.NewMessage(
									update.Message.Chat.ID,
									"Введите колличество:",
								)
								msgConfig.ReplyMarkup = amountMenu
								bot.Send(msgConfig)
								ord.State = 1
							} else if ord.State == 1 {
								ord.Name = itemCustomer.CustomerName
								ord.Telephone = itemCustomer.Phone
								ord.CustomerID = itemCustomer.ID
								ord.Amount, err = strconv.Atoi(update.Message.Text)
								fmt.Printf("index data type:    %T\n , %d\n", itemCustomer.ID, itemCustomer.ID)
								if err != nil {
									log.Println("Something went wrong: ", err.Error())
								}
								msgConfig := tgbotapi.NewMessage(
									update.Message.Chat.ID,
									"Спасибо за Ваш заказ, "+ord.Name+"! Хорошего Вам дня!")
								msgConfig.ReplyMarkup = mainMenuSecondOrder
								bot.Send(msgConfig)

								ord.DateCreate = post.GetTime()

								////////////////////////SEND post ORDER////////////////////////
								params := fmt.Sprintf(
									"{\"DateCreate\": \"%s\",\"Amount\": %d, \"CustomerID\": %d, \"ProductID\":%d}",
									ord.DateCreate,
									ord.Amount,
									ord.CustomerID,
									ord.ProductID)
								buf := bytes.NewBufferString(params)
								req, err := http.NewRequest("POST", postURL, buf)
								req.Header.Set("content-type", "application/json")
								if err != nil {
									panic(err)
								}
								resp, err := client.Do(req)
								if err != nil {
									panic(err)
								}
								itemCustomer = new(telegram.Customer)
								body, err := ioutil.ReadAll(resp.Body)
								err = json.Unmarshal(body, itemCustomer)
								if err != nil {
									fmt.Printf("Error Unmarshal")
									return
								}
								fmt.Printf("%s\n", body)
								fmt.Printf("ID: %v\n", itemCustomer.ID)
								ord.CustomerID = itemCustomer.ID
								delete(orderMap, update.Message.From.ID)
							}
							fmt.Printf("state: %v\n", ord)
						} else {
							//other messages
						}
					}

				} else {
					//--------------------------------------------------------------------------------------------------
					//---------------------------------------------NEW USER---------------------------------------------
					//--------------------------------------------------------------------------------------------------
					//esli ne komanda:
					if update.Message.Text == mainMenu.Keyboard[0][0].Text {

						orderMap[update.Message.From.ID] = new(telegram.Order)
						orderMap[update.Message.From.ID].State = 0

						msgConfig := tgbotapi.NewMessage(
							update.Message.Chat.ID,
							"Как Я могу к Вам обращаться?")
						msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						bot.Send(msgConfig)
					} else {
						ord, ok := orderMap[update.Message.From.ID]
						if ok {
							if ord.State == 0 {
								ord.Name = update.Message.Text
								msgConfig := tgbotapi.NewMessage(
									update.Message.Chat.ID,
									ord.Name+", введите пожалуйста свой номер телефона:")
								bot.Send(msgConfig)
								ord.State = 1
							} else if ord.State == 1 {
								ord.Telephone = update.Message.Text
								ord.State = 2
								msgConfig := tgbotapi.NewMessage(
									update.Message.Chat.ID,
									"Выберите пожалуйста продукт который Вы желаете заказать в магазине:")
								msgConfig.ReplyMarkup = productMenu
								bot.Send(msgConfig)
							} else if ord.State == 2 {
								ord.Product = update.Message.Text
								ord.ProductID = productId(ord.Product)
								msgConfig := tgbotapi.NewMessage(
									update.Message.Chat.ID,
									"Введите колличество:",
								)
								msgConfig.ReplyMarkup = amountMenu
								bot.Send(msgConfig)
								ord.State = 3
							} else if ord.State == 3 {
								ord.Amount, err = strconv.Atoi(update.Message.Text)
								if err != nil {
									log.Println("Something went wrong: ", err.Error())
								}
								msgConfig := tgbotapi.NewMessage(
									update.Message.Chat.ID,
									"Спасибо за Ваш заказ, "+ord.Name+"! Хорошего Вам дня!")
								msgConfig.ReplyMarkup = mainMenu
								bot.Send(msgConfig)

								////////////////////////////SEND post INSERT CUSTOMER///////////////////////
								paramsCustomer := fmt.Sprintf(
									"{\"CustomerName\": \"%s\", \"Phone\": \"%s\", \"ChatID\":%d}",
									ord.Name,
									ord.Telephone,
									update.Message.Chat.ID)

								buf1 := bytes.NewBufferString(paramsCustomer)
								reqCustomer, err := http.NewRequest("POST", postCustomerSetURL, buf1)
								reqCustomer.Header.Set("content-type", "application/json")
								if err != nil {
									panic(err)
								}
								respCustomer, err := client.Do(reqCustomer)
								if err != nil {
									panic(err)
								}
								fmt.Printf("NEW user %v\n", respCustomer)

								//////////////////////SEND get GET LAST ID CUSTOMER//////////////////////////
								reqCustomerID, err := http.NewRequest("GET", postCustomerLastIDURL, nil)
								reqCustomerID.Header.Set("content-type", "application/json")
								if err != nil {
									panic(err)
								}
								respCustomerID, err := client.Do(reqCustomerID)
								if err != nil {
									panic(err)
								}
								itemCustomer = new(telegram.Customer)
								bodyCustomer, err := ioutil.ReadAll(respCustomerID.Body)
								err = json.Unmarshal(bodyCustomer, itemCustomer)
								if err != nil {
									fmt.Printf("Error Unmarshal  GET LAST ID CUSTOMER")
									return
								}

								/////////////////////SEND post ORDER///////////////////////////
								ord.DateCreate = post.GetTime()
								params := fmt.Sprintf(
									"{\"DateCreate\": \"%s\","+
										" \"Amount\": %d, \"CustomerID\": %d, \"ProductID\":%d}",
									ord.DateCreate,
									ord.Amount,
									itemCustomer.ID,
									ord.ProductID)
								buf := bytes.NewBufferString(params)
								req, err := http.NewRequest("POST", postURL, buf)
								req.Header.Set("content-type", "application/json")
								if err != nil {
									panic(err)
								}
								resp, err := client.Do(req)
								if err != nil {
									panic(err)
								}
								itemCustomer = new(telegram.Customer)
								body, err := ioutil.ReadAll(resp.Body)
								err = json.Unmarshal(body, itemCustomer)
								if err != nil {
									fmt.Printf("Error Unmarshal")
									return
								}
								ord.CustomerID = itemCustomer.ID
								delete(orderMap, update.Message.From.ID)
							}
							fmt.Printf("state: %v\n", ord)
						} else {
							//other messages
						}
					}
				}
			}
		}
		//else if update.Message.
	}
	bot.StopReceivingUpdates()
}

func productId(product string) int {
	if product == productMenu.Keyboard[0][0].Text {
		return 1
	} else if product == productMenu.Keyboard[0][1].Text {
		return 2
	} else if product == productMenu.Keyboard[1][0].Text {
		return 3
	} else if product == productMenu.Keyboard[1][1].Text {
		return 4
	} else {
		return 0
	}
}

//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//	"io/ioutil"
//	"encoding/json"
//	"bytes"
//	"strconv"
//)
//
////tochka vhoda
//func main(){
//	botToken := "962901703:AAG2Yztqpf91ZISxm2oZElHrBUlk6hSi_IU"
//	//https://api.telegram.org/bot<token>/METHOD_NAME
//	botApi := "https://api.telegram.org/bot"
//	botUrl := botApi + botToken
//	offset := 0
//	for ;; {
//		updates, err := getUpdates(botUrl, offset)
//		if err != nil {
//			log.Println("Something went wrong: ", err.Error())
//		}
//		for _, update := range updates{
//			err = respond(botUrl, update)
//			offset = update.UpdateId + 1
//		}
//		fmt.Println(updates)
//	}
//}
//
////zapros obnavleniy
//func getUpdates(botUrl string, offset int)([]Update, error){
//
//	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//	var restResponse RestResponse
//	err = json.Unmarshal(body, &restResponse)
//	if err != nil {
//		return nil, err
//	}
//	return restResponse.Result, nil
//}
//
////otvet na obnovlenia
//func respond(botUrl string, update Update) (error){
//
//	var botMessage BotMessage
//	botMessage.ChatId = update.Message.Chat.ChatId
//	botMessage.Text = update.Message.Text
//	//buf, err := json.Marshal("Hello you")
//	buf, err := json.Marshal(botMessage)
//	if err != nil {
//		return err
//	}
//	_, err = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))
//	if err != nil {
//		return err
//	}
//	return nil
//}

//
//
//
//
//
//
//package main
//
//import (
//"accounting_software/telegram"
//"bytes"
//"encoding/json"
//"fmt"
//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
//"io/ioutil"
//"log"
//"net/http"
//"strconv"
//"time"
//)
//
//const tgbotapiKey = "962901703:AAG2Yztqpf91ZISxm2oZElHrBUlk6hSi_IU"
//
////const postURL = "http://localhost:3001/order"
//const postCustomerURL = "http://localhost:3001/customer"
//const postURL = "http://localhost:3001/order"
//
//var mainMenu = tgbotapi.NewReplyKeyboard(
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("Заказать"),
//	),
//)
//var productMenu = tgbotapi.NewReplyKeyboard(
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("🥖 батон"),
//		tgbotapi.NewKeyboardButton("🍞 кирпич"),
//	),
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("🥐 булка"),
//		tgbotapi.NewKeyboardButton("🍞 ржаной"),
//	),
//)
//var amountMenu = tgbotapi.NewReplyKeyboard(
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("1"),
//		tgbotapi.NewKeyboardButton("2"),
//	),
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("3"),
//		tgbotapi.NewKeyboardButton("4"),
//		tgbotapi.NewKeyboardButton("5"),
//	),
//)
//
//var orderMap map[int]*telegram.Order
//
////pered startom programmi
//func init() {
//	orderMap = make(map[int]*telegram.Order)
//}
//
//func main() {
//	var (
//		bot        *tgbotapi.BotAPI
//		err        error
//		updChannel tgbotapi.UpdatesChannel
//		update     tgbotapi.Update
//		updConfig  tgbotapi.UpdateConfig
//		botUser    tgbotapi.User
//
//		itemCustomer *telegram.Customer
//	)
//
//	client := http.Client{
//		Transport: &http.Transport{
//			TLSHandshakeTimeout:   time.Second * 10,
//			DisableKeepAlives:     false,
//			DisableCompression:    false,
//			IdleConnTimeout:       time.Second * 10,
//			ResponseHeaderTimeout: time.Second * 10,
//			ExpectContinueTimeout: time.Second * 10,
//		},
//		Timeout: time.Second * 10,
//	}
//
//	//inicaleziruem bota
//	bot, err = tgbotapi.NewBotAPI(tgbotapiKey)
//	if err != nil {
//		fmt.Printf("Bot init error: %s\n", err)
//		return
//	}
//
//	botUser, err = bot.GetMe()
//	if err != nil {
//		fmt.Printf("Bot getme error: %s\n", err)
//		return
//	}
//
//	fmt.Printf("auth ok! Bot is: %s\n ", botUser.FirstName)
//	//poluchili canal
//
//	updConfig.Timeout = 60
//	updConfig.Limit = 10
//	updConfig.Offset = 0
//
//	updChannel, err = bot.GetUpdatesChan(updConfig)
//	if err != nil {
//		fmt.Printf("Update channel error: %s\n", err)
//		return
//	}
//
//	//tgbotapi.InlineKeyboardMarkup{}
//	//btn := tgbotapi.KeyboardButton{Text: "Заказать товара"}
//	for {
//		update = <-updChannel
//		if update.Message != nil {
//			//esli comanda:
//
//
//			if update.Message.IsCommand() {
//				cmdText := update.Message.Command()
//				if cmdText == "start" {
//					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Здравствуйте! Для заказа нажмите кнопку заказать ⬇")
//					msg.ReplyMarkup = mainMenu
//					bot.Send(msg)
//				} else if cmdText == "menu" {
//					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню")
//					msg.ReplyMarkup = mainMenu
//					bot.Send(msg)
//				}
//			} else {
//				//esli ne komanda:
//				if update.Message.Text == mainMenu.Keyboard[0][0].Text {
//
//					orderMap[update.Message.From.ID] = new(telegram.Order)
//					orderMap[update.Message.From.ID].State = 0
//
//					//fmt.Printf(
//					//	"message: %s\n",
//					//	update.Message.Text)
//
//					msgConfig := tgbotapi.NewMessage(
//						update.Message.Chat.ID,
//						"Введите Имя:")
//					msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
//					bot.Send(msgConfig)
//
//				} else {
//					ord, ok := orderMap[update.Message.From.ID]
//					if ok {
//						if ord.State == 0 {
//							ord.Name = update.Message.Text
//							msgConfig := tgbotapi.NewMessage(
//								update.Message.Chat.ID,
//								"Введите телефон:")
//							bot.Send(msgConfig)
//							ord.State = 1
//						} else if ord.State == 1 {
//							ord.Telephone = update.Message.Text
//
//							////////////////////////////////////
//							params := fmt.Sprintf(
//								"{\"Phone\": \"%s\"}",
//								ord.Telephone)
//
//							buf := bytes.NewBufferString(params)
//							req, err := http.NewRequest("POST", postCustomerURL, buf)
//							req.Header.Set("content-type", "application/json")
//							if err != nil {
//								panic(err)
//							}
//
//							resp, err := client.Do(req)
//							if err != nil {
//								panic(err)
//							}
//							//defer resp.Body.Close()
//							itemCustomer = new(telegram.Customer)
//							body, err := ioutil.ReadAll(resp.Body)
//							err = json.Unmarshal(body, itemCustomer)
//							if err != nil {
//								fmt.Printf("Error Unmarshal")
//								return
//							}
//							fmt.Printf("%s\n", body)
//							fmt.Printf("ID: %v\n", itemCustomer.ID)
//							if itemCustomer.ID != 0{
//								ord.CustomerID = itemCustomer.ID
//							}else{
//								fmt.Printf("no user with such tel")
//
//								//params := fmt.Sprintf(
//								//	"{\"Phone\": \"%s\"}",
//								//	ord.Telephone)
//								//buf := bytes.NewBufferString(params)
//								req, err := http.NewRequest("POST", postCustomerURL, buf)
//								req.Header.Set("content-type", "application/json")
//								if err != nil {
//									panic(err)
//								}
//
//								resp, err := client.Do(req)
//								if err != nil {
//									panic(err)
//								}
//								//defer resp.Body.Close()
//								itemCustomer = new(telegram.Customer)
//								body, err := ioutil.ReadAll(resp.Body)
//								err = json.Unmarshal(body, itemCustomer)
//								if err != nil {
//									fmt.Printf("Error Unmarshal")
//									return
//								}
//								fmt.Printf("%s\n", body)
//								fmt.Printf("ID: %v\n", itemCustomer.ID)
//
//							}
//
//							fmt.Printf("index data type:    %T\n", ord.CustomerID)
//
//							////////////////////////////////////////////////
//
//							ord.State = 2
//							msgConfig := tgbotapi.NewMessage(
//								update.Message.Chat.ID,
//								"Введите название продукта:")
//							msgConfig.ReplyMarkup = productMenu
//							bot.Send(msgConfig)
//
//						} else if ord.State == 2 {
//							ord.Product = update.Message.Text
//
//							//ProductID depend from ProductName
//							if ord.Product == productMenu.Keyboard[0][0].Text {
//								ord.ProductID = 1
//							} else if ord.Product == productMenu.Keyboard[0][1].Text {
//								ord.ProductID = 2
//							} else if ord.Product == productMenu.Keyboard[1][0].Text {
//								ord.ProductID = 3
//							} else if ord.Product == productMenu.Keyboard[1][1].Text {
//								ord.ProductID = 4
//							}
//
//							msgConfig := tgbotapi.NewMessage(
//								update.Message.Chat.ID,
//								"Введите колличество:",
//							)
//							//msgConfig.ReplyMarkup = productMenu
//							msgConfig.ReplyMarkup = amountMenu
//
//							bot.Send(msgConfig)
//							ord.State = 3
//						} else if ord.State == 3 {
//							ord.Amount, err = strconv.Atoi(update.Message.Text)
//							if err != nil {
//								log.Println("Something went wrong: ", err.Error())
//							}
//							msgConfig := tgbotapi.NewMessage(
//								update.Message.Chat.ID,
//								"Спасибо за Ваш заказ!")
//							msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
//							bot.Send(msgConfig)
//
//							//fmt.Printf("state before POST: %v\n", ord)
//							//err = post.SendPost(ord)
//							//if err != nil {
//							//	fmt.Printf("Send Post error: %v\n", err)
//							//}
//
//							////////////////////////////////////////////////
//							params := fmt.Sprintf(
//								"{\"Name\": \"%s\", \"Telephone\": \"%s\", \"Product\": \"%s\","+
//									" \"Amount\": %d, \"CustomerID\": %d, \"ProductID\":%d}",
//								ord.Name,
//								ord.Telephone,
//								ord.Product,
//								ord.Amount,
//								ord.CustomerID,
//								ord.ProductID)
//
//							buf := bytes.NewBufferString(params)
//							req, err := http.NewRequest("POST", postURL, buf)
//							req.Header.Set("content-type", "application/json")
//							if err != nil {
//								panic(err)
//							}
//
//							resp, err := client.Do(req)
//							if err != nil {
//								panic(err)
//							}
//							//defer resp.Body.Close()
//							itemCustomer = new(telegram.Customer)
//							body, err := ioutil.ReadAll(resp.Body)
//							err = json.Unmarshal(body, itemCustomer)
//							if err != nil {
//								fmt.Printf("Error Unmarshal")
//								return
//							}
//							fmt.Printf("%s\n", body)
//							fmt.Printf("ID: %v\n", itemCustomer.ID)
//							ord.CustomerID = itemCustomer.ID
//							////////////////////////////////////////////
//
//							delete(orderMap, update.Message.From.ID)
//						}
//						fmt.Printf("state: %v\n", ord)
//					} else {
//						//other messages
//					}
//				}
//			}
//		}
//		//else if update.Message.
//	}
//
//	bot.StopReceivingUpdates()
//}
//
////package main
////
////import (
////	"fmt"
////	"log"
////	"net/http"
////	"io/ioutil"
////	"encoding/json"
////	"bytes"
////	"strconv"
////)
////
//////tochka vhoda
////func main(){
////	botToken := "962901703:AAG2Yztqpf91ZISxm2oZElHrBUlk6hSi_IU"
////	//https://api.telegram.org/bot<token>/METHOD_NAME
////	botApi := "https://api.telegram.org/bot"
////	botUrl := botApi + botToken
////	offset := 0
////	for ;; {
////		updates, err := getUpdates(botUrl, offset)
////		if err != nil {
////			log.Println("Something went wrong: ", err.Error())
////		}
////		for _, update := range updates{
////			err = respond(botUrl, update)
////			offset = update.UpdateId + 1
////		}
////		fmt.Println(updates)
////	}
////}
////
//////zapros obnavleniy
////func getUpdates(botUrl string, offset int)([]Update, error){
////
////	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
////	if err != nil {
////		return nil, err
////	}
////	defer resp.Body.Close()
////	body, err := ioutil.ReadAll(resp.Body)
////	if err != nil {
////		return nil, err
////	}
////	var restResponse RestResponse
////	err = json.Unmarshal(body, &restResponse)
////	if err != nil {
////		return nil, err
////	}
////	return restResponse.Result, nil
////}
////
//////otvet na obnovlenia
////func respond(botUrl string, update Update) (error){
////
////	var botMessage BotMessage
////	botMessage.ChatId = update.Message.Chat.ChatId
////	botMessage.Text = update.Message.Text
////	//buf, err := json.Marshal("Hello you")
////	buf, err := json.Marshal(botMessage)
////	if err != nil {
////		return err
////	}
////	_, err = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))
////	if err != nil {
////		return err
////	}
////	return nil
////}
