import Image from "next/image"

const Kategori = () => {
    return(
        <div className="px-4 border rounded-sm border-color-monocrome" style={{
            boxShadow: "3px 0 3px 0 rgba(0.5, 0.5, 0.5, 0.5)"
        }}>
            <div className="grid grid-cols-2 py-4">
                <Image src={'/public/assets/Banner-1.png'} height={100} width={50}/>
                <Image src={'/public/assets/Banner-2.png'} height={100} width={50}/>
                <Image src={'/public/assets/Banner-3.png'} height={100} width={50}/>
                <Image src={'/public/assets/Banner-4.png'} height={100} width={50}/>
            </div>
        </div>
    )
}

export default Kategori