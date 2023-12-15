"use client"

// import { authUserSession } from "@/lib/aut-hlibs";
import { CreditCard, ShoppingCart } from "@phosphor-icons/react";
import Image from "next/image";

const Page = () => {

    // const user = await authUserSession()

    const cardData = [
        {
            id: 1,
            title: "Paket tumpeng 1",
            imageSrc: "/assets/Banner-1-min.png",
            descrition: "Dengan 50 nasi kuning box dengan 2 jenis lauk serta 3 sayur dan air putih gelas Tambah jumlah box dengan klik di bawah setelah pemesanan",
        },
        {
            id: 2,
            title: "Paket tumpeng 2",
            imageSrc: "/assets/Banner-2-min.png",
            descrition: "Dengan 200 nasi kuning box dengan 4 jenis lauk serta 5 sayur dan air putih gelas, Tambah jumlah box dengan klik di bawah setelah pemesanan",
        },
        {
            id: 3,
            title: "Paket tumpeng 3",
            imageSrc: "/assets/Banner-3-min.png",
            descrition: "Dengan 250 nasi kuning box dengan 5 jenis lauk serta 6 sayur dan air putih gelas Tambah jumlah box dengan klik di bawah setelah pemesanan"
        }
    ]

    return (
        <>
            <h1 className="p-4 font-bold text-2xl">Tumpeng</h1>
            <div className="grid grid-cols-3 gap-4 p-4">
                {cardData.map((card) => (
                    <div key={card.id} className="p-4 max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
                        <div className="flex justify-center">
                            <Image width={300} height={200} src={card.imageSrc} alt="" />
                        </div>
                        <div className="p-5">
                                <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{card.title}</h5>                            
                            <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">{card.descrition}</p>
                            <div className="grid grid-cols-2 gap-3">
                                <button><CreditCard size={32} alt="Beli dan bayar"/></button>
                                <button><ShoppingCart size={20} alt="Masukan keranjang"/></button>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
            <h1 className="p-4 font-bold text-2xl">Nasi box</h1>
            <div className="grid grid-cols-3 gap-4 p-4">
                {cardData.map((card) => (
                    <div key={card.id} className="p-4 max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
                        <div className="flex justify-center">
                            <Image width={300} height={200} src={card.imageSrc} alt="" />
                        </div>
                        <div className="p-5">
                                <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{card.title}</h5>                            
                            <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">{card.descrition}</p>
                            <div className="grid grid-cols-2 gap-3">
                                <button><CreditCard size={32} alt="Beli dan bayar"/></button>
                                <button><ShoppingCart size={20} alt="Masukan keranjang"/></button>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
            <h1 className="p-4 font-bold text-2xl">Snack</h1>
            <div className="grid grid-cols-3 gap-4 p-4">
                {cardData.map((card) => (
                    <div key={card.id} className="p-4 max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
                        <div className="flex justify-center">
                            <Image width={300} height={200} src={card.imageSrc} alt="" />
                        </div>
                        <div className="p-5">
                                <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{card.title}</h5>                            
                            <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">{card.descrition}</p>
                            <div className="grid grid-cols-2 gap-3">
                                <button><CreditCard size={32} alt="Beli dan bayar"/></button>
                                <button><ShoppingCart size={20} alt="Masukan keranjang"/></button>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </>
    )
}

export default Page;
