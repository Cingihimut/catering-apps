import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import React from "react";
import { AddProductForm } from "./AddProductForm";
import { columns } from "./table/columns";
import { DataTable } from "./table/DataTable";

const API_URL = "http://localhost:8080/api/products";

async function fetchData(url) {
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`Failed to fetch data. Status: ${response.status}`);
  }
  return response.json();
}

export async function asyncData() {
  try {
    const data = await fetchData(API_URL);
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
          <DataTable columns={columns} data={data} />
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
