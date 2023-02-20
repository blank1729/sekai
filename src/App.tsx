import "./App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Header from "./components/Header";
import Home from "./Pages/Home";
import Account from "./Pages/Account";
import Contact from "./Pages/Contact";
import { Provider } from 'react-redux'
import { store } from "./app/store";
import Counter from "./Pages/Counter";


function App() {
  return (
    <>
      <BrowserRouter>
        <Provider store={store}>
          <Header />
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Account />} />
            <Route path="/contact" element={<Contact />} />
            <Route path="/counter" element={<Counter/>} />
          </Routes>
        </Provider>
      </BrowserRouter>
    </>
  );
}

export default App;
