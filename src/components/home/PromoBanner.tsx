import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Truck, Shield, Headphones, RefreshCw } from "lucide-react";

const features = [
  {
    icon: Truck,
    title: "Free Shipping",
    description: "On orders above ₹999",
  },
  {
    icon: Shield,
    title: "Secure Payment",
    description: "100% secure transactions",
  },
  {
    icon: Headphones,
    title: "24/7 Support",
    description: "Dedicated customer support",
  },
  {
    icon: RefreshCw,
    title: "Easy Returns",
    description: "7-day return policy",
  },
];

const PromoBanner = () => {
  return (
    <>
      {/* Design Partnership Banner */}
      <section className="py-8 gradient-primary">
        <div className="container">
          <div className="flex flex-col md:flex-row items-center justify-between gap-6 text-primary-foreground">
            <div>
              <h3 className="text-2xl font-heading font-bold mb-2">
                Design with Ease, Print with LK Printers
              </h3>
              <p className="opacity-90">
                Create stunning designs and get them printed with professional quality
              </p>
            </div>
            <Link to="/design-tools">
              <Button variant="hero" size="lg">
                Start Designing
              </Button>
            </Link>
          </div>
        </div>
      </section>

      {/* Features Bar */}
      <section className="py-8 border-b border-border">
        <div className="container">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-6">
            {features.map((feature, index) => {
              const Icon = feature.icon;
              return (
                <div key={index} className="flex items-center gap-3">
                  <div className="w-12 h-12 rounded-full bg-primary/10 flex items-center justify-center flex-shrink-0">
                    <Icon className="w-6 h-6 text-primary" />
                  </div>
                  <div>
                    <h4 className="font-semibold text-foreground text-sm">{feature.title}</h4>
                    <p className="text-xs text-muted-foreground">{feature.description}</p>
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      </section>
    </>
  );
};

export default PromoBanner;
