import { Button } from "@/components/ui/button";

import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { FoodCard } from "@/components/foodcard";

export default function Home() {
  let daftarMakanan = [
    { nama: "Makanan 1", harga: "Rp. 10K", foto: "/image/maklemak.jpg" },
    { nama: "Makanan 2", harga: "Rp. 15K", foto: "/image/soto.jpg" },
    { nama: "Makanan 3", harga: "Rp. 20K", foto: "/image/nasgor.jpg" },
  ];
  return (
    <main>
      <div className=" w-full">
        <Accordion type="single" collapsible>
          <AccordionItem value="item-1">
            <AccordionTrigger>Makan apa hari ini?</AccordionTrigger>
            <AccordionContent>Makan daging dengan sayur kol</AccordionContent>
          </AccordionItem>
        </Accordion>
      </div>
      <div className="flex flex-wrap">
        {daftarMakanan.map((item, index) => (
          <div key={index} className="w-full md:w-1/3 p-4">
            <FoodCard harga={item.harga} nama={item.nama} foto={item.foto} />
          </div>
        ))}
      </div>
      <Button>Test button</Button>
    </main>
  );
}
