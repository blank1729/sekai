import { useEffect, useState } from "react";

interface Item {
  id: string;
  name: string;
  price: string;
}

const Home = () => {
  const [items, setItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState(true);
  const getItems = async () => {
    const data = await fetch("/api/products");
    const items = await data.json();
    setItems(items);
    setLoading(false);
  };

  useEffect(() => {
    getItems();
  }, []);
  if (loading) {
    return <div>Loading</div>;
  }
  return (
    <div>
      {items.map((item) => (
        <div key={item.id}>
          <p>{item.id}</p>
          <p>{item.name}</p>
          <p>{item.price}</p>
        </div>
      ))}
    </div>
  );
};

export default Home;
