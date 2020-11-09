import React from "react";
import { Grid, IconButton, Typography, ListItem, ListItemText } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';
import { AddCircle, RemoveCircle } from '@material-ui/icons';
export const useStyles = makeStyles((theme) => ({
  img: {
    width: '150px'
  },
  autoMargin: {
    margin: 'auto'
  },
  linethrough: {
    textDecoration: 'line-through'
  },
  listItem: {
    padding: theme.spacing(1, 0),
  },
  title: {
    marginTop: theme.spacing(2),
  },
}));

export default function CartItem({ item, removeProductfromCart, addProductToCart }) {
  const classes = useStyles();
  const cartItem = item
  return (
      <Grid container className={classes.listItem} spacing={3}>
        <Grid item sm={2}>
          <IconButton aria-label="add" onClick={() => addProductToCart(cartItem.product)}>
            <AddCircle />
          </IconButton>
          <IconButton aria-label="minus" onClick={() => removeProductfromCart(cartItem.product)}>
            <RemoveCircle color="error" />
          </IconButton>
        </Grid>
        <Grid item sm={4}>
          <img alt="product image" className={classes.img} src={cartItem.product.image} />
        </Grid>
        <Grid item sm={2}>
          <Typography variant="body2" color="textSecondary">{cartItem.product.name}</Typography>
          <Typography variant="body2" color="primary">Each {convertToDollar(item.product.amount)}</Typography>
        </Grid>
        <Grid item sm={2}>
          <Typography variant="body2" color="textSecondary">Qty</Typography>
          <Typography variant="body2" color="primary">{cartItem.quantity}</Typography>
        </Grid>
        <Grid item sm={2}>
          <Typography variant="body2" color="textSecondary">Price</Typography>
          <Typography variant="body2" color="primary">{convertToDollar(cartItem.sub_total)}</Typography>
        </Grid>
      </Grid>
  )
}

function convertToDollar(amount) {
  let dollars = amount / 100;
  dollars = dollars.toLocaleString("en-US", { style: "currency", currency: "USD" });
  return dollars
}