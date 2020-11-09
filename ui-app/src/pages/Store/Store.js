import React, {useEffect} from "react";
import PropTypes from 'prop-types'
import Layout from "../Layout";
import { Products } from "./Products";
const Store = (props) => {
  const { products, loadProducts, addProductToCart} = props
  useEffect(() => {
    loadProducts()
  }, [loadProducts]);
  return (
    <Layout title="STORE">
      <Products products={products} addProductToCart={addProductToCart} />
    </Layout>
  );
}

Store.propTypes = {
  products: PropTypes.array,
  loadProducts: PropTypes.func
}
Store.defaultProps = {
  products: [],
  loadProducts: () => {}
}

export default Store;