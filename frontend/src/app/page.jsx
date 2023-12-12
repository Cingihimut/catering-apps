import Footer from "@/components/Footer/index";
import NavBar from "@/components/NavBar/index";
import Kategori from "./kategori/page";
import Carousel from "./bannerPromo/caraousel";

const Page = () => {
  return (
    <>
      <NavBar />
      <main>
        <div className="py-6">
          <Carousel />
        </div>
      </main>
      <p className="font-bold p-4">Kategori</p>
      <Kategori />
      <Footer />
    </>
  );
};
export default Page;
