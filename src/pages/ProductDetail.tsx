import { useState } from "react";
import { useParams, Link } from "react-router-dom";
import Layout from "@/components/layout/Layout";
import { Button } from "@/components/ui/button";
import { Star, ShoppingCart, Heart, Share2, Truck, Shield, RefreshCw, Minus, Plus } from "lucide-react";
import { useToast } from "@/hooks/use-toast";

const ProductDetail = () => {
  const { id } = useParams();
  const { toast } = useToast();
  const [quantity, setQuantity] = useState(1);
  const [selectedVariant, setSelectedVariant] = useState("standard");

  // Mock product data
  const product = {
    id,
    name: "Premium Visiting Cards",
    price: 199,
    originalPrice: 299,
    rating: 4.8,
    reviews: 1250,
    description: "Make a lasting impression with our premium quality visiting cards. Printed on high-quality 350 GSM paper with a smooth matte finish.",
    features: [
      "Premium 350 GSM paper",
      "Matte or glossy finish",
      "Full color printing",
      "Quick turnaround time",
      "Free design templates",
    ],
    variants: [
      { id: "standard", name: "Standard", price: 199 },
      { id: "premium", name: "Premium", price: 299 },
      { id: "luxury", name: "Luxury", price: 499 },
    ],
    images: ["/placeholder.svg", "/placeholder.svg", "/placeholder.svg"],
  };

  const handleAddToCart = () => {
    toast({
      title: "Added to cart",
      description: `${quantity}x ${product.name} added to your cart`,
    });
  };

  return (
    <Layout>
      <div className="container py-8">
        <nav className="text-sm text-muted-foreground mb-6">
          <Link to="/" className="hover:text-primary">Home</Link>
          <span className="mx-2">/</span>
          <Link to="/products" className="hover:text-primary">Products</Link>
          <span className="mx-2">/</span>
          <span className="text-foreground">{product.name}</span>
        </nav>

        <div className="grid md:grid-cols-2 gap-8 lg:gap-12">
          {/* Product Images */}
          <div className="space-y-4">
            <div className="aspect-square bg-muted rounded-xl overflow-hidden">
              <img
                src={product.images[0]}
                alt={product.name}
                className="w-full h-full object-cover"
              />
            </div>
            <div className="grid grid-cols-3 gap-4">
              {product.images.map((img, idx) => (
                <div key={idx} className="aspect-square bg-muted rounded-lg overflow-hidden cursor-pointer hover:ring-2 ring-primary">
                  <img src={img} alt="" className="w-full h-full object-cover" />
                </div>
              ))}
            </div>
          </div>

          {/* Product Info */}
          <div>
            <h1 className="text-2xl md:text-3xl font-heading font-bold mb-4">{product.name}</h1>
            
            <div className="flex items-center gap-4 mb-4">
              <div className="flex items-center gap-1">
                {Array.from({ length: 5 }).map((_, i) => (
                  <Star key={i} className={`w-4 h-4 ${i < Math.floor(product.rating) ? "fill-amber-400 text-amber-400" : "text-muted"}`} />
                ))}
              </div>
              <span className="text-sm text-muted-foreground">
                {product.rating} ({product.reviews} reviews)
              </span>
            </div>

            <div className="flex items-baseline gap-3 mb-6">
              <span className="text-3xl font-bold text-foreground">₹{product.variants.find(v => v.id === selectedVariant)?.price}</span>
              <span className="text-lg text-muted-foreground line-through">₹{product.originalPrice}</span>
              <span className="bg-success/10 text-success px-2 py-1 rounded-md text-sm font-medium">
                {Math.round((1 - product.price / product.originalPrice) * 100)}% OFF
              </span>
            </div>

            <p className="text-muted-foreground mb-6">{product.description}</p>

            {/* Variants */}
            <div className="mb-6">
              <h3 className="font-semibold mb-3">Select Type</h3>
              <div className="flex gap-3">
                {product.variants.map((variant) => (
                  <button
                    key={variant.id}
                    onClick={() => setSelectedVariant(variant.id)}
                    className={`px-4 py-2 rounded-lg border transition-colors ${
                      selectedVariant === variant.id
                        ? "border-primary bg-primary/10 text-primary"
                        : "border-border hover:border-primary/50"
                    }`}
                  >
                    {variant.name}
                  </button>
                ))}
              </div>
            </div>

            {/* Quantity */}
            <div className="mb-6">
              <h3 className="font-semibold mb-3">Quantity</h3>
              <div className="flex items-center gap-3">
                <button
                  onClick={() => setQuantity(Math.max(1, quantity - 1))}
                  className="w-10 h-10 rounded-lg border border-border flex items-center justify-center hover:bg-muted"
                >
                  <Minus className="w-4 h-4" />
                </button>
                <span className="w-12 text-center font-medium">{quantity}</span>
                <button
                  onClick={() => setQuantity(quantity + 1)}
                  className="w-10 h-10 rounded-lg border border-border flex items-center justify-center hover:bg-muted"
                >
                  <Plus className="w-4 h-4" />
                </button>
              </div>
            </div>

            {/* Actions */}
            <div className="flex gap-4 mb-8">
              <Button variant="accent" size="lg" className="flex-1" onClick={handleAddToCart}>
                <ShoppingCart className="w-5 h-5 mr-2" />
                Add to Cart
              </Button>
              <Button variant="outline" size="lg">
                <Heart className="w-5 h-5" />
              </Button>
              <Button variant="outline" size="lg">
                <Share2 className="w-5 h-5" />
              </Button>
            </div>

            {/* Features */}
            <div className="border-t border-border pt-6">
              <h3 className="font-semibold mb-4">Features</h3>
              <ul className="space-y-2">
                {product.features.map((feature, idx) => (
                  <li key={idx} className="flex items-center gap-2 text-sm text-muted-foreground">
                    <div className="w-1.5 h-1.5 rounded-full bg-primary" />
                    {feature}
                  </li>
                ))}
              </ul>
            </div>

            {/* Trust Badges */}
            <div className="grid grid-cols-3 gap-4 mt-6 pt-6 border-t border-border">
              <div className="text-center">
                <Truck className="w-6 h-6 mx-auto text-primary mb-1" />
                <span className="text-xs text-muted-foreground">Fast Delivery</span>
              </div>
              <div className="text-center">
                <Shield className="w-6 h-6 mx-auto text-primary mb-1" />
                <span className="text-xs text-muted-foreground">Quality Assured</span>
              </div>
              <div className="text-center">
                <RefreshCw className="w-6 h-6 mx-auto text-primary mb-1" />
                <span className="text-xs text-muted-foreground">Easy Returns</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default ProductDetail;
