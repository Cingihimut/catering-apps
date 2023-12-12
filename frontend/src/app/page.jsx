import Footer from "@/components/Footer/index";
import NavBar from "@/components/NavBar/index";
import Kategori from "./kategori/page";
import Carousel from "./bannerPromo/caraousel";

const Page = () => {
  return (
    <>
<<<<<<< HEAD
      <NavBar/>
        <main>
          <div className="py-6">
            <Carousel/>
          </div>
        </main>
        <p className="font-bold p-4 text-2xl">Kategori</p>
        <Kategori/>
      <Footer/>
=======
      <NavBar />
      <main>
        <div className="py-6">
          <Carousel />
        </div>
      </main>
      <p className="font-bold p-4">Kategori</p>
      <Kategori />
      <Footer />
>>>>>>> d03c1138372929ba3c229c7a75033a8e22a9f073
    </>
  );
};
export default Page;
