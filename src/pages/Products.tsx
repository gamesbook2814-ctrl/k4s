import { useState } from "react";
import Layout from "@/components/layout/Layout";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Star, Filter, Grid, List, ChevronDown } from "lucide-react";
import { Link } from "react-router-dom";

const allProducts = [
  { id: 1, name: "Premium Visiting Cards", price: 199, originalPrice: 299, rating: 4.8, reviews: 1250, category: "stationery", image: "/placeholder.svg" },
  { id: 2, name: "Round Neck T-Shirts", price: 349, originalPrice: 499, rating: 4.7, reviews: 890, category: "apparel", image: "/placeholder.svg" },
  { id: 3, name: "Custom Photo Mugs", price: 249, originalPrice: 399, rating: 4.9, reviews: 2100, category: "photo-gifts", image: "/placeholder.svg" },
  { id: 4, name: "Sticker Labels", price: 149, originalPrice: 199, rating: 4.6, reviews: 680, category: "packaging", image: "/placeholder.svg" },
  { id: 5, name: "Desktop Calendar 2026", price: 299, originalPrice: 449, rating: 4.8, reviews: 450, category: "calendars", image: "/placeholder.svg" },
  { id: 6, name: "Premium Notebook", price: 399, originalPrice: 599, rating: 4.7, reviews: 320, category: "stationery", image: "/placeholder.svg" },
  { id: 7, name: "Polo T-Shirts", price: 449, originalPrice: 649, rating: 4.6, reviews: 560, category: "apparel", image: "/placeholder.svg" },
  { id: 8, name: "Custom Hoodies", price: 899, originalPrice: 1299, rating: 4.8, reviews: 340, category: "apparel", image: "/placeholder.svg" },
  { id: 9, name: "Photo Frames", price: 349, originalPrice: 499, rating: 4.7, reviews: 780, category: "photo-gifts", image: "/placeholder.svg" },
  { id: 10, name: "Brochures", price: 199, originalPrice: 299, rating: 4.5, reviews: 420, category: "marketing", image: "/placeholder.svg" },
  { id: 11, name: "Banners", price: 599, originalPrice: 899, rating: 4.6, reviews: 290, category: "marketing", image: "/placeholder.svg" },
  { id: 12, name: "Tote Bags", price: 249, originalPrice: 399, rating: 4.8, reviews: 650, category: "packaging", image: "/placeholder.svg" },
];

const categories = [
  { id: "all", name: "All Products" },
  { id: "stationery", name: "Business Stationery" },
  { id: "apparel", name: "Apparel" },
  { id: "photo-gifts", name: "Photo Gifts" },
  { id: "packaging", name: "Packaging" },
  { id: "marketing", name: "Marketing & Promo" },
  { id: "calendars", name: "Calendars & Diaries" },
];

const Products = () => {
  const [selectedCategory, setSelectedCategory] = useState("all");
  const [searchQuery, setSearchQuery] = useState("");
  const [viewMode, setViewMode] = useState<"grid" | "list">("grid");

  const filteredProducts = allProducts.filter((product) => {
    const matchesCategory = selectedCategory === "all" || product.category === selectedCategory;
    const matchesSearch = product.name.toLowerCase().includes(searchQuery.toLowerCase());
    return matchesCategory && matchesSearch;
  });

  return (
    <Layout>
      <div className="container py-8">
        {/* Breadcrumb */}
        <nav className="text-sm text-muted-foreground mb-6">
          <Link to="/" className="hover:text-primary">Home</Link>
          <span className="mx-2">/</span>
          <span className="text-foreground">All Products</span>
        </nav>

        <div className="flex flex-col lg:flex-row gap-8">
          {/* Sidebar Filters */}
          <aside className="w-full lg:w-64 flex-shrink-0">
            <div className="bg-card rounded-xl border border-border p-4">
              <h3 className="font-heading font-semibold mb-4 flex items-center gap-2">
                <Filter className="w-4 h-4" /> Categories
              </h3>
              <ul className="space-y-2">
                {categories.map((category) => (
                  <li key={category.id}>
                    <button
                      onClick={() => setSelectedCategory(category.id)}
                      className={`w-full text-left px-3 py-2 rounded-lg text-sm transition-colors ${
                        selectedCategory === category.id
                          ? "bg-primary text-primary-foreground"
                          : "hover:bg-muted"
                      }`}
                    >
                      {category.name}
                    </button>
                  </li>
                ))}
              </ul>
            </div>
          </aside>

          {/* Products Grid */}
          <div className="flex-1">
            {/* Toolbar */}
            <div className="flex flex-col sm:flex-row gap-4 justify-between items-start sm:items-center mb-6">
              <div className="flex-1 max-w-md">
                <Input
                  placeholder="Search products..."
                  value={searchQuery}
                  onChange={(e) => setSearchQuery(e.target.value)}
                />
              </div>
              <div className="flex items-center gap-4">
                <span className="text-sm text-muted-foreground">
                  {filteredProducts.length} products
                </span>
                <div className="flex border border-border rounded-lg overflow-hidden">
                  <button
                    onClick={() => setViewMode("grid")}
                    className={`p-2 ${viewMode === "grid" ? "bg-primary text-primary-foreground" : "hover:bg-muted"}`}
                  >
                    <Grid className="w-4 h-4" />
                  </button>
                  <button
                    onClick={() => setViewMode("list")}
                    className={`p-2 ${viewMode === "list" ? "bg-primary text-primary-foreground" : "hover:bg-muted"}`}
                  >
                    <List className="w-4 h-4" />
                  </button>
                </div>
              </div>
            </div>

            {/* Products */}
            <div className={viewMode === "grid" 
              ? "grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4"
              : "flex flex-col gap-4"
            }>
              {filteredProducts.map((product) => (
                <Link
                  key={product.id}
                  to={`/product/${product.id}`}
                  className={`bg-card rounded-xl border border-border overflow-hidden hover:shadow-lg transition-all group ${
                    viewMode === "list" ? "flex" : ""
                  }`}
                >
                  <div className={`relative ${viewMode === "list" ? "w-32 h-32" : "aspect-square"} bg-muted overflow-hidden`}>
                    <img
                      src={product.image}
                      alt={product.name}
                      className="w-full h-full object-cover group-hover:scale-105 transition-transform"
                    />
                  </div>
                  <div className="p-4 flex-1">
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
        </div>
      </div>
    </Layout>
  );
};

export default Products;
