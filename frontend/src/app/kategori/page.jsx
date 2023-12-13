import Image from "next/image";

const Kategori = () => {
  return (
    <div
      className="p-4 border rounded-xl"
      style={{
        boxShadow: "3px 0 3px 0 rgba(0.5, 0.5, 0.5, 0.5)",
      }}
    >
      <div className="grid grid-cols-2 p-4">
        <Image
          src={"/assets/Banner-1-min.png"}
          height={100}
          width={50}
          alt="Kategori 1"
        />
        <Image
          src={"/assets/Banner-2-min.png"}
          height={100}
          width={50}
          alt="Kategori 2"
        />
        <Image
          src={"/assets/Banner-3-min.png"}
          height={100}
          width={50}
          alt="Kategori 3"
        />
        <Image
          src={"/assets/Banner-4-min.png"}
          height={100}
          width={50}
          alt="Kategori 4"
        />
      </div>
    </div>
  );
};

export default Kategori;
