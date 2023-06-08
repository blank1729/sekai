import {
  createBrowserRouter,
  createRoutesFromElements,
  Outlet,
  Route,
} from "react-router-dom";
import Cart from "../pages/cart/cart";
import Home from "../pages/home/home";
import Header from "../components/Header/Header";

export const router = createBrowserRouter(
  createRoutesFromElements(
    <Route
      element={
        <>
          <Header />
          <Outlet />
        </>
      }
    >
      <Route path="/" element={<Home />} />
      <Route path="/cart" element={<Cart />} />
    </Route>
  )
);
