// app/transaksi/TransactionForm.jsx
"use client";
import React, { useState, useEffect } from "react";

const TransactionForm = ({ selectedMenuItem }) => {
  const [address, setAddress] = useState("Bojonegoro no surender");
  const [selectedDate, setSelectedDate] = useState("");
  const [quantity, setQuantity] = useState(1);
  const [selectedItem, setSelectedItem] = useState(null);

  useEffect(() => {
    const currentDate = new Date().toISOString().split("T")[0];
    setSelectedDate(currentDate);
  }, []);

  const handleDateChange = (e) => {
    setSelectedDate(e.target.value);
  };

  const handleQuantityChange = (e) => {
    setQuantity(e.target.value);
  };

  const handleOrderNow = () => {
    console.log("Pesanan berhasil!", {
      address,
      selectedDate,
      quantity,
      selectedItem,
    });
  };

  useEffect(() => {
    setSelectedItem(selectedMenuItem);
  }, [selectedMenuItem]);

  return (
    <div>
      <div>
        <label>Alamat:</label>
        <input
          type="text"
          value={address}
          onChange={(e) => setAddress(e.target.value)}
        />
      </div>
      <div>
        <label>Tanggal:</label>
        <input type="date" value={selectedDate} onChange={handleDateChange} />
      </div>

      {selectedItem && (
        <div>
          <h2>{selectedItem.name}</h2>
          <img src={`gambar${selectedItem.name}.png`} alt={selectedItem.name} />
          <p>Harga: ${selectedItem.price}</p>
        </div>
      )}

      <div>
        <label>Jumlah:</label>
        <input type="number" value={quantity} onChange={handleQuantityChange} />
      </div>

      <button
        className="max-w-xs mx-auto bg-blue-500 mt-24 my-24  text-white"
        onClick={handleOrderNow}
      >
        Pesan Sekarang
      </button>
    </div>
  );
};

export default TransactionForm;
