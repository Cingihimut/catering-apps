import Link from "next/link"
import InputSearch from "./InputSearch"

const NavBar = () => {
    return (
        <header>
                <div className="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl p-4 gap-2">
                    <Link href="/" className="flex items-center space-x-3 rtl:space-x-reverse">
                        <span className="self-center text-2xl font-semibold whitespace-nowrap text-color-third dark:text-white">My Kitchen</span>
                    </Link>
                    <div className="flex items-center space-x-6 rtl:space-x-reverse">
                        <InputSearch/>
                        <Link href="/login" className="text-sm text-color-dark font-bold dark:text-blue-500 hover:underline">Login</Link>
                    </div>
                </div>
                <div className="flex justify-normal px-4 gap-4">
                    <Link href="/menu/rekomendasi">Rekomendasi kami</Link>
                    <Link href="/menu/terbaik">Pilihan Terbaik</Link>
                    <Link href="/menu/terlaris">Pilihan Terlaris</Link>
                    <Link href="/menu">Menu lengkap</Link>
                </div>
        </header>
    )
}

export default NavBar