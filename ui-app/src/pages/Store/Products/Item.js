import React from "react";
import { Button, Grid, Card, CardHeader, CardMedia, CardActions, Typography, CardContent } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  root: {
    maxWidth: 345,
    margin: "auto"
  },
  media: {
    height: 0,
    paddingTop: "56.25%" // 16:9
  },
  alignFlexEnd:{
    justifyContent: 'flex-end'
  }
}));

export function ProductItem({ product, addProductToCart }) {
  const classes = useStyles();
  const addToCartHandler = () => {
    addProductToCart(product)
  }
  var dollars = product.amount / 100;
  dollars = dollars.toLocaleString("en-US", {style:"currency", currency:"USD"});
  return (
    <Grid item xs={12} sm={3} >
      <Card className={classes.root}>
        <CardHeader title={product.name} />
        <CardMedia
          className={classes.media}
          image={product.image}
          title="Orange"
        />
        <CardContent>
          <Typography variant="body2" color="textSecondary" component="p">
            {product.details}
          </Typography>
          <Typography>
            {dollars}
          </Typography>
        </CardContent>
        <CardActions className={classes.alignFlexEnd}>
          <Button variant="contained" color="primary" onClick={addToCartHandler}>
            Add to cart
          </Button>
        </CardActions>
      </Card>
    </Grid>
  )
}