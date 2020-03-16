import Products from './views/Products'
import ProductsUpdate from './views/ProductsUpdate'
import ProductAdd from './views/ProductAdd'
import Decommissioned from './views/Decommissioned'
import Suppliers from './views/Suppliers'
import SuppliersUpdate from './views/SuppliersUpdate'
import SupplierAdd from './views/SupplierAdd'
import Categories from './views/Categories'
import CategoriesUpdate from './views/CategoriesUpdate'
import CategoryAdd from './views/CategoryAdd'
import test from './components/Chart'
import CashBox from "./views/CashBox";
import Supply from "./views/Supply";
import SupplyAdd from "./views/SupplyAdd";
import SupplyUpdate from "./views/SupplyUpdate";
import Report from "./views/Report";
import DayReport from "./views/DayReport";
import Orders from "./views/Orders";
// import UtilityBills from "./views/UtilityBills";
import NewUtilityBills from "./views/NewUtilityBill";
import Welcome from "./views/Welcome";

export default [
    {
        path: '/',
        name: 'welcome',
        component: Welcome,
    },
    {
        path: '/products',
        name: 'products',
        component: Products,
    },
    {
        path: '/products/update',
        name: 'productsUpdate',
        component: ProductsUpdate,
    },
    {
        path: '/products/add',
        name: 'productAdd',
        component: ProductAdd,
    },
    {
        path: '/supply/decommissioned',
        name: 'decommissioned',
        component: Decommissioned
    },
    {
        path: '/suppliers',
        name: 'suppliers',
        component: Suppliers
    },
    {
        path: '/suppliers/update',
        name: 'suppliersUpdate',
        component: SuppliersUpdate
    },
    {
        path: '/suppliers/add',
        name: 'supplierAdd',
        component: SupplierAdd
    },
    {
        path: '/categories',
        name: 'categories',
        component: Categories
    },
    {
        path: '/categories/update',
        name: 'categoriesUpdate',
        component: CategoriesUpdate
    },
    {
        path: '/categories/add',
        name: 'categoryAdd',
        component: CategoryAdd
    },
    {
        path: '/test',
        name: 'test',
        component: test
    },
    {
        path: '/cashbox',
        name: 'cashBox',
        component: CashBox
    },

    {
        path: '/supply',
        name: 'supply',
        component: Supply
    },

    {
        path: '/supply/add',
        name: 'supplyAdd',
        component: SupplyAdd
    },

    {
        path: '/supply/update',
        name: 'supplyUpdate',
        component: SupplyUpdate
    },
    {
        path: '/report',
        name: 'report',
        component: Report
    },
    {
        path: '/day_report',
        name: 'dayReport',
        component: DayReport
    },
    {
        path: '/utility_bills',
        name: 'newUtilityBills',
        component: NewUtilityBills
    },
    {
        path: '/orders',
        name: 'orders',
        component: Orders
    },
]
