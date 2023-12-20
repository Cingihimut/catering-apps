import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import React from "react";
import { AddProductForm } from "./AddProductForm";
import { ProductsTable } from "./table/ProductsTable";

const API_URL = process.env.API_URL + "/api/products";
async function fetchData() {
  const response = await fetch(API_URL);
  if (!response.ok) {
    throw new Error(`Failed to fetch data. Status: ${response.status}`);
  }
  return response.json();
}

export async function asyncData() {
  try {
    const data = await fetchData();
    return data.data;
  } catch (error) {
    console.error(error.message);
    return { data: [] };
  }
}

const DashboardTabs = async () => {
  const data = await asyncData();

  return (
    <Tabs defaultValue="daftar" className="flex flex-col items-center">
      <TabsList className="flex space-x-8">
        <TabsTrigger value="daftar">Daftar Produk</TabsTrigger>
        <TabsTrigger value="tambah">Tambah Produk</TabsTrigger>
      </TabsList>
      <TabsContent value="daftar">
        <div className="container w-full">
          <ProductsTable data={data} />
        </div>
      </TabsContent>
      <TabsContent value="tambah">
        <div className="container w-full">
          <AddProductForm />
        </div>
      </TabsContent>
    </Tabs>
  );
};

export default DashboardTabs;
