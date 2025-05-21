import "./Products.css";
import { useEffect, useState } from "react";

interface Product {
  id: number;
  name: string;
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
}
interface ProductsProps {
  onSelectionChange: (quantities: Record<number, number>) => void;
  onProductList: (products: Product[]) => void;
}
function Products({ onSelectionChange, onProductList }: ProductsProps) {
  const [products, setProducts] = useState<Product[]>([]);
  const [quantities, setQuantities] = useState<Record<number, number>>({});

  useEffect(() => {
    fetch("http://localhost:8080/products")
      .then((res) => res.json())
      .then((data: Product[]) => {
        setProducts(data);
        onProductList(data);
      })
      .catch((err) => console.error("Błąd pobierania produktów:", err));
  }, []);

  const handleChange = (id: number, value: string) => {
    const num = parseFloat(value);
    const newQuantities = {
      ...quantities,
      [id]: isNaN(num) ? 0 : num,
    };
    setQuantities(newQuantities);
    onSelectionChange(newQuantities); // aktualizuj nadrzędny komponent
  };

  return (
    <div className="products-frame">
      <div className="products-panel">
        <Heading />
        <ul className="products-list">
          {products.map((p) => {
            const quantity = quantities[p.id] || 0;
            return (
              <li
                key={p.id}
                className={`product-item ${quantity > 0 ? "selected" : ""}`}
              >
                <div className="product-info">
                  <strong>{p.name}</strong>
                  <span>
                    {p.calories} kcal, {p.protein}g białka, {p.fat}g tłuszczu,{" "}
                    {p.carbs}g węgli
                  </span>
                </div>
                <input
                  type="number"
                  min="0"
                  placeholder="g"
                  value={quantity}
                  onChange={(e) => handleChange(p.id, e.target.value)}
                  className="product-input"
                />
              </li>
            );
          })}
        </ul>
      </div>
    </div>
  );
}

function Heading() {
  return <div className="products-heading">Wybierz produkty</div>;
}

export default Products;
