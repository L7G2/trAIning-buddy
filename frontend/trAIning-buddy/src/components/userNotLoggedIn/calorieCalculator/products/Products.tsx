import "./Products.css";

function Products() {
  return (
    <div className="products-frame">
      <ProductsPanel />
    </div>
  );
}
function ProductsPanel() {
  return (
    <div className="products-panel">
      <Heading />
    </div>
  );
}
function Heading() {
  return <div className="products-heading">Wybierz produkty</div>;
}

export default Products;
