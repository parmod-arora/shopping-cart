import React from 'react';
import PropTypes from 'prop-types'
import {
  Badge,
  AppBar,
  Toolbar,
  IconButton,
  List,
  ListItem,
  ListItemText,
  Typography, Link as MLink
} from "@material-ui/core";
import { Link } from "react-router-dom";
import { Home } from "@material-ui/icons";
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(theme => ({
  navbarDisplayFlex: {
    display: `flex`,
    justifyContent: `space-between`
  },
  title: {
    flexGrow: 1,
  },
  navDisplayFlex: {
    display: `flex`,
    justifyContent: `space-between`
  },
  linkText: {
    textDecoration: `none`,
    cursor: 'pointer',
    textTransform: `uppercase`,
    color: theme.palette.header.main,
    '&:hover': {
      textDecoration: `none`,
    },
  },
  badge: {
    padding: theme.spacing(0, 1)
  },
  icon: {
    padding: theme.spacing(0, 1)
  }
}));

export const Header = ({logout, cartItemsCount}) => {
  const classes = useStyles();
  return (
    <AppBar position="static">
      <Toolbar>
          <IconButton edge="start" color="inherit" aria-label="home">
            <Home fontSize="large" />
          </IconButton>
          <Typography variant="h6" className={classes.title}>
            <Link className={classes.linkText} to={'/products'} key={'home'}>
              Fruit Shop
            </Link>
          </Typography>
          <List component="nav" className={classes.navDisplayFlex}>
            <Link to={'/products'} key={'product'} className={classes.linkText}>
                <ListItem><ListItemText primary={'product'} /></ListItem>
            </Link>
            <Link to={'/cart'} key={'cart'} className={classes.linkText}>
              <ListItem>
                {cartItemsCount && <IconButton className={classes.icon}>
                  <Badge badgeContent={cartItemsCount} color="secondary">
                    <ShoppingCartIcon />
                  </Badge>
                </IconButton>}
                <ListItemText primary={'cart'} />
              </ListItem>
            </Link>
            <MLink onClick={() => logout()} className={classes.linkText}>
              <ListItem>
                <ListItemText primary={'logout'} className={classes.linkText} />
              </ListItem>
            </MLink>
          </List>
      </Toolbar>
    </AppBar>
  );
};

Header.defaultProps = {
  logout: () =>{},
  cartItemsCount: 0
}
Header.prototype = {
  logout: PropTypes.func,
  cartItemsCount: PropTypes.number
}

export default Header