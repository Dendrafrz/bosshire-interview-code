import React, { useState } from 'react';

const Product = () => {
  const [quantity, setQuantity] = useState(1);

  const addToCart = async () => {
    try {
      const response = await fetch('/api/cart', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ quantity }),
      });
      const data = await response.json();
      console.log(data);
      // You can show a success message or redirect the user to the cart page
    } catch (error) {
      console.error('Error adding to cart:', error);
      // You can show an error message to the user
    }
  };

  return (
    <div>
      <h1>Product Details</h1>
      <h2>Product Name</h2>
      <p>Product Description</p>
      <p>Price: $10</p>
      <p>
        Quantity:
        <input
          type="number"
          value={quantity}
          onChange={(e) => setQuantity(Number(e.target.value))}
        />
      </p>
      <button onClick={addToCart}>Add to Cart</button>
    </div>
  );
};

export default Product;