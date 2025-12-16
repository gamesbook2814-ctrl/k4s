import { MapPin, Package } from "lucide-react";
import { Link } from "react-router-dom";

const TopBar = () => {
  return (
    <div className="bg-primary text-primary-foreground text-sm">
      <div className="container flex items-center justify-between py-2">
        <Link 
          to="/same-day-delivery" 
          className="hover:underline flex items-center gap-1"
        >
          4 Hrs Express Delivery now available in major cities
        </Link>
        <div className="hidden md:flex items-center gap-6">
          <Link to="/track-order" className="hover:underline flex items-center gap-1">
            <Package className="w-4 h-4" />
            Track Order
          </Link>
          <Link to="/store-locator" className="hover:underline flex items-center gap-1">
            <MapPin className="w-4 h-4" />
            Store Locator
          </Link>
          <Link to="/sample-kit" className="hover:underline">
            Sample Kit
          </Link>
          <Link 
            to="/business-solutions" 
            className="bg-accent text-accent-foreground px-3 py-1 rounded-md hover:bg-accent-dark transition-colors"
          >
            Business Solutions
          </Link>
        </div>
      </div>
    </div>
  );
};

export default TopBar;
