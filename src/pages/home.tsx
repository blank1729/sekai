import { useEffect, useState } from "react";

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
    <main>
      <section className="container">
        <h1>Items</h1>
        <div
          style={{
            display: "flex",
            gap: "1rem",
          }}
        >
          {items.map((item) => (
            <div key={item.id}>
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
