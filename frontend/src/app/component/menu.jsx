"use client";
import DaftarMenu from "./daftar/DaftarMenu";
import Image from "next/image";
import { useState, useEffect } from "react";
export default function Menu() {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(`https://labs.mhdaris.me/api/products/`);
        if (response.ok) {
          const result = await response.json();
          setData(result.data);
        } else {
          setError(`Error fetching food data: ${response.status}`);
        }
      } catch (error) {
        setError(`Error fetching food data: ${error.message}`);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  return (
    <section>
      <div className="absolute left-0  right-0 w-full justify-start">
        <div className="h-48 w-48 absolute -top-[180px] -left-12 text-left -z-10">
          <Image
            src={"/selada-kiri.png"}
            layout={"fill"}
            objectFit={"contain"}
            alt={"salad"}
          />
        </div>
        <div className="h-48 w-48 absolute -top-[200px] -right-12 text-right -z-10">
          <Image
            src={"/selada-kanan.png"}
            layout={"fill"}
            objectFit={"contain"}
            alt={"salad"}
          />
        </div>
      </div>
      <div className="text-center font-semibold color text-4xl">
        <h1 className="">Menu</h1>
      </div>
      <div className="grid grid-cols-3 gap-4 mt-12">
        {data.map((produk) => (
          <DaftarMenu product={produk} />
        ))}
      </div>
    </section>
  );
}
