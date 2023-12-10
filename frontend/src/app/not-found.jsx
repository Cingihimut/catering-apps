"use client"

import { FileSearch } from "@phosphor-icons/react";
import { useRouter } from "next/navigation";


const Page = () => {

    const router = useRouter()

    return(
    <div className="min-h-screen max-w--xl mx-auto flex justify-center items-center">
        <div className="flex justify-center items-center gap-4">
            <FileSearch size={40} className="text-color-dark"/>
            <h3 className="text-color-dark transition-all text-2xl font-bold">Page Not Found</h3>
            <button onClick={() => router.back()} className="text-color-dark hover:text-color-third">Kembali</button>
        </div>
    </div>
)
}

export default Page;