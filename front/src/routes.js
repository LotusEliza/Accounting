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
import test from './views/test'
import CashBox from "./views/CashBox";
// categoryAdd

export default [
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
        path: '/products/decommissioned',
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


]
