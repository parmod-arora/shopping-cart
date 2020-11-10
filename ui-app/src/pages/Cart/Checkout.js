import React, { useEffect } from "react";
import PropTypes from 'prop-types'
import Layout from "../Layout";
import { Grid, Button } from "@material-ui/core";
import CartItem from "./CartItem";
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import { useFormik } from 'formik';
import { makeStyles } from '@material-ui/core/styles';
import { InputText } from "../../component/Input/InputText";
export const useStyles = makeStyles((theme) => ({
  container: {
    marginTop: theme.spacing(2)
  },
  listItem: {
    padding: theme.spacing(1, 2),
  },
  title: {
    marginTop: theme.spacing(2),
    padding: theme.spacing(1, 0),
  },
  applyCouponForm: {
    marginLeft: theme.spacing(2),
    padding: theme.spacing(2, 0),
    display: 'flex',
    width: '100%',
  },
  formButton: {
    marginLeft: theme.spacing(1),
    flex: 1,
    lineHeight: '1.187'
  },
  formInput: {
    flex: 2
  },
  alignRight: {
    textAlign: "right"
  },
  alignCenter: {
    textAlign: "center"
  }
}));

function Checkout(props) {
  const { lineItems, cartItems, fetchUserCart, addProductToCart, removeProductfromCart, subTotalAmount, totalSavingAmount, totalAmount, dispatchApplyCoupon, dispatchRemoveCoupon, coupon = {}, dispatchCheckout } = props
  const classes = useStyles();
  useEffect(() => {
    fetchUserCart()
  }, [fetchUserCart]);

  const formik = useFormik({
    initialValues: {
      coupon: '',
    },
    validate,
    validateOnBlur: false,
    validateOnChange: true,
    onSubmit: values => {
      dispatchApplyCoupon({
        coupon: values.coupon
      })
    },
  })

  if (cartItems.length <= 0) {
    return (<Layout>
      <Grid className={classes.container} container spacing={3}>
        <Grid item sm={12} className={classes.alignCenter}>
          <Typography variant="h4">Your cart is empty</Typography>
        </Grid>
      </Grid>
    </Layout>)
  }

  const couponName = coupon.name
  return (
    <Layout>

      <Grid className={classes.container} container spacing={3}>
        <Grid item sm={12} className={classes.title}><Typography variant="h6">Products</Typography></Grid>
        <Grid item sm={8}>
          {
            cartItems.map(item => <CartItem
              className={classes.listItem}
              addProductToCart={addProductToCart}
              removeProductfromCart={removeProductfromCart}
              key={item.id} item={item} />)
          }
        </Grid>
        <Grid item sm={4} >
          <Grid container spacing={3}>
            <Grid item sm={8}>
              <Typography variant="h6">Sub Total</Typography>
            </Grid>
            <Grid item sm={4} className={classes.alignRight}>
              {subTotalAmount && <Typography variant="h6">{convertToDollar(subTotalAmount)}</Typography>}
            </Grid>
          </Grid>
          <List disablePadding>
            {
              lineItems.map((lineItem) => {
                const discount = lineItem.discount_applied
                return (
                  <Grid key={discount.name} container spacing={3}>
                    <Grid item sm={5}>
                      <Typography variant="body2">{discount.name}</Typography>
                    </Grid>
                    <Grid item sm={4}>
                      <Typography variant="body2">{`${discount.discount} %`}</Typography>
                    </Grid>
                    <Grid item sm={3} className={classes.alignRight}>
                      <Typography variant="body2">-{convertToDollar(lineItem.discount_amount)}</Typography>
                    </Grid>
                  </Grid>
                )
              })
            }
            {totalSavingAmount &&
              <Grid container spacing={3}>
                <Grid item sm={8}><Typography variant="h6">Total Saving</Typography></Grid>
                <Grid className={classes.alignRight} item sm={4}><Typography variant="h6">{convertToDollar(totalSavingAmount)}</Typography></Grid>
              </Grid>
            }
            {totalAmount &&
              <Grid container spacing={3}>
                <Grid item sm={8}><Typography variant="h6">Total Amount</Typography></Grid>
                <Grid className={classes.alignRight} item sm={4}><Typography variant="h6">{convertToDollar(totalAmount)}</Typography></Grid>
              </Grid>
            }
            {!couponName &&
              <Grid container spacing={3}>
                <form className={classes.applyCouponForm} onSubmit={formik.handleSubmit} noValidate autoComplete="off">
                  <InputText variant="outlined" className={classes.formInput} name="coupon" margin="none" id="coupon" label="Use COUPON_30" onChange={formik.handleChange} onBlur={formik.handleBlur} value={formik.values.coupon} errorMsg={formik.touched.coupon && formik.errors.coupon} />
                  <Button variant="contained" color="primary" type="submit" className={classes.formButton}>Apply Coupon</Button>
                </form>
              </Grid>
            }
            {couponName &&
              <Grid container spacing={3}>
                <form className={classes.applyCouponForm} onSubmit={e => {
                  e.preventDefault()
                  dispatchRemoveCoupon({
                    id: coupon.id
                  })
                }} autoComplete="off">
                  <InputText variant="outlined" className={classes.formInput} margin="none" disabled value={coupon.name} />
                  <Button variant="contained" color="secondary" type="submit" className={classes.formButton}>Remove Coupon</Button>
                </form>
              </Grid>
            }
            <Grid container spacing={3}>
              <Button variant="contained" color="primary" onClick={dispatchCheckout} className={classes.formButton}>Checkout</Button>
            </Grid>
          </List>
        </Grid>
      </Grid>
    </Layout>
  )
}
function convertToDollar(amount) {
  let dollars = amount / 100;
  dollars = dollars.toLocaleString("en-US", { style: "currency", currency: "USD" });
  return dollars
}

const validate = values => {
  const errors = {};
  if (!values.coupon) {
    errors.coupon = 'Required';
  }
  return errors;
};

const noop = () => { }
Checkout.defaultProps = {
  lineItems: [],
  cartItems: [],
  fetchUserCart: noop,
  addProductToCart: noop,
  removeProductfromCart: noop,
  subTotalAmount: 0,
  totalSavingAmount: 0,
  totalAmount: 0,
  dispatchApplyCoupon: noop,
  dispatchRemoveCoupon: noop,
  coupon: {},
  dispatchCheckout: noop
}
Checkout.prototype = {
  lineItems: PropTypes.array,
  cartItems: PropTypes.array,
  fetchUserCart: PropTypes.func,
  addProductToCart: PropTypes.func,
  removeProductfromCart: PropTypes.func,
  subTotalAmount: PropTypes.number,
  totalSavingAmount: PropTypes.number,
  totalAmount: PropTypes.number,
  dispatchApplyCoupon: PropTypes.func,
  dispatchRemoveCoupon: PropTypes.func,
  coupon: PropTypes.object,
  dispatchCheckout: PropTypes.func,
}

export default Checkout