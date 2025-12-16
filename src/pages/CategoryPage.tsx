import { useParams } from "react-router-dom";
import Layout from "@/components/layout/Layout";
import { Link } from "react-router-dom";
import { Star } from "lucide-react";

const categoryData: Record<string, { title: string; description: string; products: any[] }> = {
  "calendars-diaries": {
    title: "Calendars & Diaries",
    description: "Premium quality calendars and diaries for every occasion",
    products: [
      { id: 1, name: "Desktop Calendar 2026", price: 299, originalPrice: 449, rating: 4.8, reviews: 450, image: "/placeholder.svg" },
      { id: 2, name: "Wall Calendar 2026", price: 399, originalPrice: 599, rating: 4.7, reviews: 320, image: "/placeholder.svg" },
      { id: 3, name: "Executive Diary", price: 599, originalPrice: 899, rating: 4.9, reviews: 210, image: "/placeholder.svg" },
      { id: 4, name: "Pocket Diary", price: 199, originalPrice: 299, rating: 4.6, reviews: 180, image: "/placeholder.svg" },
    ],
  },
  "apparel": {
    title: "Custom Apparel",
    description: "High quality custom printed t-shirts, hoodies, and more",
    products: [
      { id: 1, name: "Round Neck T-Shirts", price: 349, originalPrice: 499, rating: 4.7, reviews: 890, image: "/placeholder.svg" },
      { id: 2, name: "Polo T-Shirts", price: 449, originalPrice: 649, rating: 4.6, reviews: 560, image: "/placeholder.svg" },
      { id: 3, name: "Custom Hoodies", price: 899, originalPrice: 1299, rating: 4.8, reviews: 340, image: "/placeholder.svg" },
      { id: 4, name: "Custom Caps", price: 199, originalPrice: 299, rating: 4.5, reviews: 280, image: "/placeholder.svg" },
    ],
  },
  "packaging": {
    title: "Packaging Solutions",
    description: "Custom boxes, labels, stickers and packaging materials",
    products: [
      { id: 1, name: "Flat Mailer Boxes", price: 499, originalPrice: 749, rating: 4.7, reviews: 320, image: "/placeholder.svg" },
      { id: 2, name: "Sticker Labels", price: 149, originalPrice: 199, rating: 4.6, reviews: 680, image: "/placeholder.svg" },
      { id: 3, name: "Tote Bags", price: 249, originalPrice: 399, rating: 4.8, reviews: 650, image: "/placeholder.svg" },
      { id: 4, name: "Custom Labels", price: 199, originalPrice: 299, rating: 4.5, reviews: 420, image: "/placeholder.svg" },
    ],
  },
  "stationery": {
    title: "Business Stationery",
    description: "Professional business cards, letterheads, and office essentials",
    products: [
      { id: 1, name: "Premium Visiting Cards", price: 199, originalPrice: 299, rating: 4.8, reviews: 1250, image: "/placeholder.svg" },
      { id: 2, name: "Letterheads", price: 299, originalPrice: 449, rating: 4.6, reviews: 380, image: "/placeholder.svg" },
      { id: 3, name: "Premium Notebook", price: 399, originalPrice: 599, rating: 4.7, reviews: 320, image: "/placeholder.svg" },
      { id: 4, name: "Envelopes", price: 149, originalPrice: 199, rating: 4.5, reviews: 240, image: "/placeholder.svg" },
    ],
  },
  "corporate-gifts": {
    title: "Corporate Gifts",
    description: "Premium gifting solutions for your business partners and employees",
    products: [
      { id: 1, name: "Welcome Kit", price: 1499, originalPrice: 2199, rating: 4.9, reviews: 180, image: "/placeholder.svg" },
      { id: 2, name: "Premium Mug Set", price: 599, originalPrice: 899, rating: 4.7, reviews: 320, image: "/placeholder.svg" },
      { id: 3, name: "Corporate Backpack", price: 899, originalPrice: 1299, rating: 4.8, reviews: 150, image: "/placeholder.svg" },
      { id: 4, name: "Awards & Trophies", price: 799, originalPrice: 1199, rating: 4.6, reviews: 210, image: "/placeholder.svg" },
    ],
  },
  "photo-gifts": {
    title: "Photo Gifts",
    description: "Personalized photo frames, mugs, prints and more",
    products: [
      { id: 1, name: "Custom Photo Mugs", price: 249, originalPrice: 399, rating: 4.9, reviews: 2100, image: "/placeholder.svg" },
      { id: 2, name: "Photo Frames", price: 349, originalPrice: 499, rating: 4.7, reviews: 780, image: "/placeholder.svg" },
      { id: 3, name: "Photo Prints", price: 99, originalPrice: 149, rating: 4.8, reviews: 1500, image: "/placeholder.svg" },
      { id: 4, name: "Photo Books", price: 599, originalPrice: 899, rating: 4.6, reviews: 420, image: "/placeholder.svg" },
    ],
  },
  "drinkware": {
    title: "Drinkware",
    description: "Premium mugs, flasks, and sipper bottles",
    products: [
      { id: 1, name: "Premium Mugs", price: 299, originalPrice: 449, rating: 4.7, reviews: 520, image: "/placeholder.svg" },
      { id: 2, name: "Premium Flasks", price: 599, originalPrice: 899, rating: 4.8, reviews: 320, image: "/placeholder.svg" },
      { id: 3, name: "Sipper Bottles", price: 399, originalPrice: 599, rating: 4.6, reviews: 410, image: "/placeholder.svg" },
      { id: 4, name: "Travel Mugs", price: 349, originalPrice: 499, rating: 4.5, reviews: 280, image: "/placeholder.svg" },
    ],
  },
};

const CategoryPage = () => {
  const { category } = useParams();
  const data = categoryData[category || ""] || {
    title: "Products",
    description: "Browse our products",
    products: [],
  };

  return (
    <Layout>
      <div className="container py-8">
        <nav className="text-sm text-muted-foreground mb-6">
          <Link to="/" className="hover:text-primary">Home</Link>
          <span className="mx-2">/</span>
          <span className="text-foreground">{data.title}</span>
        </nav>

        <div className="mb-8">
          <h1 className="text-3xl font-heading font-bold mb-2">{data.title}</h1>
          <p className="text-muted-foreground">{data.description}</p>
        </div>

        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          {data.products.map((product) => (
            <Link
              key={product.id}
              to={`/product/${product.id}`}
              className="bg-card rounded-xl border border-border overflow-hidden hover:shadow-lg transition-all group"
            >
              <div className="aspect-square bg-muted overflow-hidden">
                <img
                  src={product.image}
                  alt={product.name}
                  className="w-full h-full object-cover group-hover:scale-105 transition-transform"
                />
              </div>
              <div className="p-4">
                <h3 className="font-medium text-foreground line-clamp-2 mb-2">{product.name}</h3>
                <div className="flex items-center gap-1 mb-2">
                  <Star className="w-3 h-3 fill-amber-400 text-amber-400" />
                  <span className="text-xs text-muted-foreground">
                    {product.rating} ({product.reviews})
                  </span>
                </div>
                <div className="flex items-center gap-2">
                  <span className="font-semibold text-foreground">₹{product.price}</span>
                  <span className="text-sm text-muted-foreground line-through">₹{product.originalPrice}</span>
                </div>
              </div>
            </Link>
          ))}
        </div>
      </div>
    </Layout>
  );
};

export default CategoryPage;
