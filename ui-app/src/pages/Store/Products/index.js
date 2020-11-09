import React from "react";
import { makeStyles } from '@material-ui/core/styles';
import { Grid } from '@material-ui/core';
import { ProductItem } from './Item';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    marginTop: theme.spacing(8)
  }
}));

export function Products({products, addProductToCart}) {
  const classes = useStyles();
  return <div className={classes.root}>
  <Grid container spacing={3}>
    {
      products.map(p => <ProductItem key={p.id} product={p} addProductToCart={addProductToCart} />)
    }
  </Grid>
</div>
}