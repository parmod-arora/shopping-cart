import * as React from "react";
import { Container, Typography, Grid } from "@material-ui/core";
import Header  from "./Header";
export const Layout = ({ title, children }) => {
  return (
    <React.Fragment>
      <Header />
      <Container maxWidth="lg">
        {/* Title */}
        <Grid container justify="center" direction="row">
          <Typography variant="h4">{title}</Typography>
        </Grid>
        <Grid container >
          {children}
        </Grid>
      </Container>
    </React.Fragment>
  );
};
export default Layout;
