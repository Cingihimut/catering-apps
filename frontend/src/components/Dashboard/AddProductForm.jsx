"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import * as z from "zod";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";

const formSchema = z.object({
  product_name: z.string().min(2, {
    message: "Product name must be at least 2 characters.",
  }),
  description: z.string().min(5, {
    message: "Description must be at least 5 characters.",
  }),
  price: z.number().positive({
    message: "Price must be a positive number.",
  }),
  categories: z.array(z.string()).min(1, {
    message: "Please select at least one category.",
  }),
  images: z.array(z.string()).min(1, {
    message: "Please upload at least one image.",
  }),
});

export function AddProductForm() {
  const form = useForm({
    resolver: zodResolver(formSchema),
    mode: "onChange",
  });

  const onSubmit = (data) => {
    // Handle form submission logic here
    console.log(data);
  };

  return (
    <Form form={form} className="p-0 m-0">
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-3">
        <FormField
          control={form.control}
          name="product_name"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="text-primary">Product Name</FormLabel>
              <FormControl>
                <Input placeholder="Product Name" {...field} />
              </FormControl>

              <FormMessage />
            </FormItem>
          )}
        />

        {/* Description */}
        <FormField
          control={form.control}
          name="description"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="text-primary">Description</FormLabel>
              <FormControl>
                <Textarea
                  placeholder="Product Description"
                  className="resize-none"
                  {...field}
                />
              </FormControl>

              <FormMessage />
            </FormItem>
          )}
        />

        {/* Price */}
        <FormField
          control={form.control}
          name="price"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="text-primary">Price</FormLabel>
              <FormControl>
                <Input type="number" placeholder="Price" {...field} />
              </FormControl>

              <FormMessage />
            </FormItem>
          )}
        />

        {/* Categories */}
        <FormField
          control={form.control}
          name="categories"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="text-primary">Categories</FormLabel>
              <FormControl>
                <Input type="text" placeholder="Categories" {...field} />
              </FormControl>

              <FormMessage />
            </FormItem>
          )}
        />

        {/* Images */}
        <FormField
          control={form.control}
          name="images"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="text-primary">Images</FormLabel>
              <FormControl>
                <Input
                  accept=".jpg, .jpeg, .png, .svg, .gif"
                  type="file"
                  onChange={(e) =>
                    field.onChange(e.target.files ? e.target.files[0] : null)
                  }
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <Button type="submit">Tambah</Button>
      </form>
    </Form>
  );
}
