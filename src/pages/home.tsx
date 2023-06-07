import { useEffect, useState } from "react";
import "./home.scss";

interface Item {
  id: string;
  name: string;
  price: string;
}

const Home = () => {
  const [items, setItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState(true);
  const path = "https://sekaii.vercel.app/api/products";

  const getItems = async () => {
    const data = await fetch(path);
    const items = await data.json();
    setItems(items);
    setLoading(false);
  };

  useEffect(() => {
    getItems();
  }, []);
  if (loading) {
    return (
      <main>
        <div className="container">Loading</div>
      </main>
    );
  }
  return (
    <main className="homepage">
      <section className="container">
        <h1>Items</h1>
        <hr />
        <div className="items">
          {items.map((item) => (
            <div key={item.id} className="item">
              <p>{item.id}</p>
              <p>{item.name}</p>
              <p>{item.price}</p>
            </div>
          ))}
        </div>
      </section>
    </main>
  );
};

export default Home;
