import { Link } from "react-router-dom";
import { Star, ArrowRight } from "lucide-react";
import { Button } from "@/components/ui/button";

const products = [
  {
    id: 1,
    name: "Premium Visiting Cards",
    price: "₹199",
    originalPrice: "₹299",
    image: "/placeholder.svg",
    rating: 4.8,
    reviews: 1250,
    href: "/products/visiting-cards",
    badge: "Best Seller",
  },
  {
    id: 2,
    name: "Round Neck T-Shirts",
    price: "₹349",
    originalPrice: "₹499",
    image: "/placeholder.svg",
    rating: 4.7,
    reviews: 890,
    href: "/products/round-neck-tshirts",
    badge: null,
  },
  {
    id: 3,
    name: "Custom Photo Mugs",
    price: "₹249",
    originalPrice: "₹399",
    image: "/placeholder.svg",
    rating: 4.9,
    reviews: 2100,
    href: "/products/mugs",
    badge: "Popular",
  },
  {
    id: 4,
    name: "Sticker Labels",
    price: "₹149",
    originalPrice: "₹199",
    image: "/placeholder.svg",
    rating: 4.6,
    reviews: 680,
    href: "/products/stickers",
    badge: null,
  },
  {
    id: 5,
    name: "Desktop Calendar 2026",
    price: "₹299",
    originalPrice: "₹449",
    image: "/placeholder.svg",
    rating: 4.8,
    reviews: 450,
    href: "/products/calendars",
    badge: "New",
  },
  {
    id: 6,
    name: "Premium Notebook",
    price: "₹399",
    originalPrice: "₹599",
    image: "/placeholder.svg",
    rating: 4.7,
    reviews: 320,
    href: "/products/notebooks",
    badge: null,
  },
];

const BestSellers = () => {
  return (
    <section className="py-12 md:py-16 bg-muted">
      <div className="container">
        <div className="flex justify-between items-center mb-10">
          <div>
            <h2 className="text-2xl md:text-3xl font-heading font-bold">
              Best Sellers
            </h2>
            <p className="text-muted-foreground mt-1">
              Most loved products by our customers
            </p>
          </div>
          <Link to="/products">
            <Button variant="outline" className="hidden sm:flex">
              View All <ArrowRight className="w-4 h-4 ml-1" />
            </Button>
          </Link>
        </div>

        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4 md:gap-6">
          {products.map((product) => (
            <Link
              key={product.id}
              to={product.href}
              className="bg-background rounded-xl overflow-hidden border border-border hover:border-primary/50 hover:shadow-lg transition-all group"
            >
              <div className="relative aspect-square bg-muted overflow-hidden">
                <img
                  src={product.image}
                  alt={product.name}
                  className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
                />
                {product.badge && (
                  <span className={`absolute top-2 left-2 text-xs font-medium px-2 py-1 rounded-md ${
                    product.badge === "Best Seller" ? "bg-accent text-accent-foreground" :
                    product.badge === "New" ? "bg-success text-success-foreground" :
                    "bg-primary text-primary-foreground"
                  }`}>
                    {product.badge}
                  </span>
                )}
              </div>
              <div className="p-3">
                <h3 className="font-medium text-sm text-foreground line-clamp-2 mb-2">
                  {product.name}
                </h3>
                <div className="flex items-center gap-1 mb-2">
                  <Star className="w-3 h-3 fill-amber-400 text-amber-400" />
                  <span className="text-xs text-muted-foreground">
                    {product.rating} ({product.reviews})
                  </span>
                </div>
                <div className="flex items-center gap-2">
                  <span className="font-semibold text-foreground">{product.price}</span>
                  <span className="text-xs text-muted-foreground line-through">{product.originalPrice}</span>
                </div>
              </div>
            </Link>
          ))}
        </div>

        <div className="mt-8 text-center sm:hidden">
          <Link to="/products">
            <Button variant="outline">
              View All Products <ArrowRight className="w-4 h-4 ml-1" />
            </Button>
          </Link>
        </div>
      </div>
    </section>
  );
};

export default BestSellers;
