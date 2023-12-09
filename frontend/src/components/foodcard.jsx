import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import Image from "next/image";

export function FoodCard({ nama, harga, foto }) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>{nama}</CardTitle>
        <Image src={foto} width={300} height={300} />
        <CardDescription>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Maxime,
          dolores!
        </CardDescription>
      </CardHeader>
      <CardContent className="grid gap-6">Harga : {harga}</CardContent>
      <CardFooter>
        <Button className="w-full">Continue</Button>
      </CardFooter>
    </Card>
  );
}
