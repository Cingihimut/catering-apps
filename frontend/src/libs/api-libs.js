export const getProductResponse = async(resources, query) => {
    const response = await fetch(`${process.env.NEXT_PUBLIC_BASE_URL}/${resources}?${query}`);
    const product = await response.json();
    return product
}