import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

export function ProductsTable({ data }) {
  return (
    <Table>
      <TableCaption>Daftar Produk</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">No</TableHead>
          <TableHead>Nama Produk</TableHead>
          <TableHead>Description</TableHead>
          <TableHead>Price</TableHead>
          <TableHead className="text-right">Categories</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {data.map((produk, index) => {
          let no = index + 1;
          return (
            <TableRow key={produk.id}>
              <TableCell className="font-medium">{no}</TableCell>
              <TableCell>{produk.product_name}</TableCell>
              <TableCell>{produk.description}</TableCell>
              <TableCell>{produk.price}</TableCell>
              <TableCell className="text-right">
                {produk.categories.map((cat) => cat.name).join(",")}
              </TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}
