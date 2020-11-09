import { Router, Switch, Route } from "react-router-dom";
import { history } from "./store/store";
import SignIn from "./pages/SignIn";
import SignUp from "./pages/Signup";
import Loading from "./pages/Common/Loading";
import Message from "./pages/Common/Message";
import Products from "./pages/Store";
import Cart from "./pages/Cart";
import { ProtectedRoute } from "./component/ProtectedRoute";

function App() {
  return (
    <Router history={history}>
      <Loading />
      <Message />
      <Switch>
        <ProtectedRoute exact path="/" component={Products} />
        <Route path="/login"><SignIn /></Route>
        <Route path="/signup"><SignUp /></Route>
        <ProtectedRoute path="/products" component={Products} />
        <ProtectedRoute path="/cart" component={Cart} />
      </Switch>
    </Router>
  );
}

export default App;
