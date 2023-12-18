import DaftarMenu from "./daftar/DaftarMenu";
import Image from "next/image";

export default function Menu() {
  let test = [1, 2, 3, 4, 5, 6];
  return (
    <section>
      <div className="absolute left-0  right-0 w-full justify-start">
        <div className="h-48 w-48 absolute -top-[180px]  -left-12 text-left -z-10">
          <Image
            src={"/selada-kiri.png"}
            layout={"fill"}
            objectFit={"contain"}
            alt={"salad"}
          />
        </div>
        <div className="h-48 w-48 absolute -top-[200px] -right-12 text-right -z-10">
          <Image
            src={"/selada-kanan.png"}
            layout={"fill"}
            objectFit={"contain"}
            alt={"salad"}
          />
        </div>
      </div>
      <div className="text-center font-semibold color text-4xl">
        <h1 className="">Menu</h1>
      </div>
      <div className="grid grid-cols-3 gap-4 mt-12">
        {test.map((el) => (
          <DaftarMenu name={el} key={el} />
        ))}
      </div>
    </section>
  );
}
