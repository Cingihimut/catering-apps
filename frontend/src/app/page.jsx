import Footer from "@/components/Footer/index"
import NavBar from "@/components/NavBar/index"
import Kategori from "./kategori/page"
import Carousel from "./bannerPromo/page"

const Page = () => {

  return (
    <>
      <NavBar/>
        <main className="App">
          <div className="max-w-lg">
            <Carousel>
              
            </Carousel>
          </div>

        </main>
        <Kategori/>
      <Footer/>
    </>
  )
}
export default Page