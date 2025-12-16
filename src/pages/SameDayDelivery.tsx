import Layout from "@/components/layout/Layout";
import { Button } from "@/components/ui/button";
import { Clock, Truck, Check, MapPin } from "lucide-react";
import { Link } from "react-router-dom";

const sameDayProducts = [
  { id: 1, name: "Visiting Cards", price: 249, image: "/placeholder.svg", time: "4 hrs" },
  { id: 2, name: "Letterheads", price: 399, image: "/placeholder.svg", time: "4 hrs" },
  { id: 3, name: "Flyers", price: 299, image: "/placeholder.svg", time: "4 hrs" },
  { id: 4, name: "Posters", price: 499, image: "/placeholder.svg", time: "4 hrs" },
  { id: 5, name: "Stickers", price: 199, image: "/placeholder.svg", time: "4 hrs" },
  { id: 6, name: "Labels", price: 149, image: "/placeholder.svg", time: "4 hrs" },
];

const cities = ["Bangalore", "Hyderabad", "Chennai", "Delhi", "Mumbai", "Pune"];

const SameDayDelivery = () => {
  return (
    <Layout>
      {/* Hero */}
      <div className="gradient-accent py-16 text-accent-foreground">
        <div className="container text-center">
          <div className="flex items-center justify-center gap-2 mb-4">
            <Clock className="w-8 h-8" />
            <span className="text-2xl font-bold">4 Hours Express</span>
          </div>
          <h1 className="text-4xl md:text-5xl font-heading font-bold mb-4">
            Same Day Delivery
          </h1>
          <p className="text-xl opacity-90 max-w-2xl mx-auto mb-8">
            Get your prints delivered within 4 hours! Available in select cities.
          </p>
          <div className="flex flex-wrap justify-center gap-3">
            {cities.map((city) => (
              <span key={city} className="bg-background/20 px-4 py-2 rounded-full flex items-center gap-1">
                <MapPin className="w-4 h-4" /> {city}
              </span>
            ))}
          </div>
        </div>
      </div>

      {/* How it works */}
      <div className="container py-12">
        <h2 className="text-2xl font-heading font-bold text-center mb-10">How It Works</h2>
        <div className="grid md:grid-cols-3 gap-8 max-w-4xl mx-auto">
          <div className="text-center">
            <div className="w-16 h-16 rounded-full bg-primary/10 flex items-center justify-center mx-auto mb-4">
              <span className="text-2xl font-bold text-primary">1</span>
            </div>
            <h3 className="font-semibold mb-2">Order Before 12 PM</h3>
            <p className="text-muted-foreground text-sm">Place your order before noon for same-day delivery</p>
          </div>
          <div className="text-center">
            <div className="w-16 h-16 rounded-full bg-primary/10 flex items-center justify-center mx-auto mb-4">
              <span className="text-2xl font-bold text-primary">2</span>
            </div>
            <h3 className="font-semibold mb-2">We Print Express</h3>
            <p className="text-muted-foreground text-sm">Your order gets priority in our production queue</p>
          </div>
          <div className="text-center">
            <div className="w-16 h-16 rounded-full bg-primary/10 flex items-center justify-center mx-auto mb-4">
              <span className="text-2xl font-bold text-primary">3</span>
            </div>
            <h3 className="font-semibold mb-2">Delivered in 4 Hours</h3>
            <p className="text-muted-foreground text-sm">Receive your prints at your doorstep</p>
          </div>
        </div>
      </div>

      {/* Products */}
      <div className="bg-muted py-12">
        <div className="container">
          <h2 className="text-2xl font-heading font-bold text-center mb-10">Available Products</h2>
          <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
            {sameDayProducts.map((product) => (
              <Link
                key={product.id}
                to={`/product/${product.id}`}
                className="bg-background rounded-xl p-4 text-center hover:shadow-lg transition-shadow"
              >
                <div className="relative">
                  <img src={product.image} alt={product.name} className="w-full aspect-square object-cover rounded-lg bg-muted mb-3" />
                  <span className="absolute top-2 right-2 bg-accent text-accent-foreground text-xs px-2 py-1 rounded-full font-medium">
                    {product.time}
                  </span>
                </div>
                <h3 className="font-medium text-sm mb-1">{product.name}</h3>
                <p className="text-primary font-semibold">₹{product.price}</p>
              </Link>
            ))}
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default SameDayDelivery;
