import Image from "next/image";

const Kategori = () => {


  const productImg = [
    {
      id: 1,
      title: "Paket 1",
      imageSrc: "/assets/Banner-1-min.png"
    },
    {
      id: 2,
      title: "Paket 2",
      imageSrc: "/assets/Banner-2-min.png"
    },
    {
      id: 3,
      title: "Paket 3",
      imageSrc: "/assets/Banner-3-min.png"
    },
    {
      id: 4,
      title: "Paket 4",
      imageSrc: "/assets/Banner-4-min.png"
    }
  ]

  return (
    <>
      <div
      className="p-4 border rounded-xl"
      style={{
        boxShadow: "3px 0 3px 0 rgba(0.5, 0.5, 0.5, 0.5)",
      }}
    >
      <div className="flex justify-between">
        <Image
          src={"/assets/Banner-1-min.png"}
          height={400}
          width={200}
          alt="Kategori 1"
          className="p-1"
        />
        <Image
          src={"/assets/Banner-2-min.png"}
          height={400}
          width={200}
          alt="Kategori 2"
          className="p-1"
        />
        <Image
          src={"/assets/Banner-3-min.png"}
          height={400}
          width={200}
          alt="Kategori 2"
          className="p-1"
        />
        <Image
          src={"/assets/Banner-4-min.png"}
          height={400}
          width={200}
          alt="Kategori 2"
          className="p-1"
        />
      </div>
    </div>
    </>
  );
};

export default Kategori;
