import Banner from "./bannerPromo/page"
import Footer from "@/components/Footer/index"
import NavBar from "@/components/NavBar/index"
import Kategori from "./kategori/page"

const Page = () => {
  return (
    <>
      <NavBar/>
      <Banner/>
      <div className="px-6 font-bold text-xl">
        Kategori
        <Kategori/>
      </div>
      <Footer/>
    </>
  )
}
export default Page