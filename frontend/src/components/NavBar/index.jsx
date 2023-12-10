import Link from "next/link"
import InputSearch from "./InputSearch"

const NavBar = () => {
    return (
        <header>
                <div className="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl p-4">
                    <Link href="/" className="flex items-center space-x-3 rtl:space-x-reverse">
                        <span className="self-center text-2xl font-semibold whitespace-nowrap">My Kitchen</span>
                    </Link>
                    <div className="flex items-center space-x-6 rtl:space-x-reverse">
                        <InputSearch/>
                        <Link href="/login" className="text-sm text-color-dark font-bold dark:text-blue-500 hover:underline">Login</Link>
                    </div>
                </div>
                <div className="flex justify-normal px-4 gap-4 text-color-third">
                    <Link href="/menu/rekomendasi/page" className="hover:text-color-dark transition duration-300 ease-in-out">Rekomendasi kami</Link>
                    <Link href="/menu/terbaik/page" className="hover:text-color-dark transition duration-300 ease-in-out">Pilihan Terbaik</Link>
                    <Link href="/menu/terlaris/page" className="hover:text-color-dark transition duration-300 ease-in-out">Pilihan Terlaris</Link>
                    <Link href="/menu" className="hover:text-color-dark transition duration-300 ease-in-out">Menu lengkap</Link>
                </div>
        </header>
    )
}

export default NavBar