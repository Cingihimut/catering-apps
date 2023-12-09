<<<<<<< HEAD
import Footer from "@/components/Footer"
import NavBar from "@/components/NavBar"
import Link from "next/link"


const Page = () => {
  return (
    <>
      <NavBar />
      <section className="bg-white dark:bg-gray-900">
        <div className="py-8 px-4 mx-auto max-w-screen-xl text-center lg:py-16 lg:px-12">
          <Link href="#" className="inline-flex justify-between items-center py-1 px-1 pr-4 mb-7 text-sm text-gray-700 bg-gray-100 rounded-full dark:bg-gray-800 dark:text-white hover:bg-gray-200 dark:hover:bg-gray-700" role="alert">
            <span className="text-l font-bold bg-color-dark rounded-full text-white px-4 py-1.5 mr-3">Baru</span> <span className="text-sm font-medium">Promo terbaru dari kami</span>
            <svg className="ml-2 w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd"></path></svg>
          </Link>
          <h1 className="mb-4 text-4xl font-extrabold tracking-tight leading-none text-gray-900 md:text-5xl lg:text-6xl dark:text-white">We invest in the world’s potential</h1>
          <p className="mb-8 text-lg font-normal text-gray-500 lg:text-xl sm:px-16 xl:px-48 dark:text-gray-400">Here at Flowbite we focus on markets where technology, innovation, and capital can unlock long-term value and drive economic growth.</p>
        </div>
        <Link href="/compone">
          Pilihan terbaik
        </Link>
      </section>
      <Footer />
    </>
  )
}
export default Page
=======
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
            <AccordionContent>Makan daging dengan sayur kangkung</AccordionContent>
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
>>>>>>> 0ca57cf31d383e4aa317ee058a7882b7967ad705
