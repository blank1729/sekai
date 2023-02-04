import "./App.css";
import Header from "./components/Header";
import Footer from "./components/Footer";
import MenuItem from "./components/MenuItem";
import { item } from "../types/item";
import Search from "./components/Search";

let items: Array<item> = [
  {
    id: "carrot",
    name: "a crazy carrot",
    price: 15,
    img: "carrot.svg",
  },
  {
    id: "pizza",
    name: "a papparoni pizza",
    price: 300,
    img: "pizza.svg",
  },
  {
    id: "corn",
    name: "a corny corn",
    price: 40,
    img: "corn.svg",
  },
  {
    id: "meatball",
    name: "a spicy meatball",
    price: 15,
    img: "meatball.svg",
  },
];

function App() {
  return (
    <>
      <Header />
      <main className="ml-[4.5rem] lg:px-[10%] pt-4">
        {/* seconday nav */}
        <div className="px-16 flex items-center">
          <div>
            <h1 className="text-2xl font-bold">SEKAI</h1>
            <p className=" flex gap-2 group">
              <p className="group-hover:underline">pragathi nagar, Bangalore</p>{" "}
              <img className="w-4" src="/arrow-down.svg" alt="" />{" "}
            </p>
          </div>
          <Search/>
        </div>
        <section className="menu mt-8">
          <div className="grid gap-4 grid-cols-2 md:grid-cols-2 lg:grid-cols-4">
            {items.map((item) => (
              <MenuItem
                name={item.name}
                img={item.img}
                price={item.price}
                key={item.id}
                id={item.id}
              />
            ))}
          </div>
        </section>
      </main>
      {/* <Footer/> */}
    </>
  );
}

export default App;
