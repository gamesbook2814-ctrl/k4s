import { Search, HelpCircle, User, ShoppingCart } from "lucide-react";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";

const Header = () => {
  const [searchQuery, setSearchQuery] = useState("");
  const [cartCount] = useState(0);

  return (
    <header className="bg-background border-b border-border sticky top-0 z-40">
      <div className="container flex items-center justify-between py-4 gap-8">
        {/* Logo */}
        <Link to="/" className="flex-shrink-0">
          <h1 className="text-2xl md:text-3xl font-heading font-bold">
            <span className="text-primary">LK</span>
            <span className="text-foreground"> Printers</span>
          </h1>
        </Link>

        {/* Search Bar */}
        <div className="flex-1 max-w-xl hidden md:block">
          <div className="relative">
            <Input
              type="search"
              placeholder="Search for products..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="pr-10 h-11 rounded-full border-border focus:border-primary"
            />
            <Button
              size="icon"
              variant="ghost"
              className="absolute right-1 top-1/2 -translate-y-1/2 h-9 w-9 rounded-full hover:bg-primary/10"
            >
              <Search className="w-5 h-5 text-muted-foreground" />
            </Button>
          </div>
        </div>

        {/* Right Actions */}
        <div className="flex items-center gap-2 md:gap-4">
          <Link to="/help" className="hidden sm:flex items-center gap-2 text-sm text-muted-foreground hover:text-foreground transition-colors">
            <HelpCircle className="w-5 h-5" />
            <span className="hidden lg:inline">Help Center</span>
          </Link>

          <Link to="/login" className="flex items-center gap-2 text-sm text-muted-foreground hover:text-foreground transition-colors">
            <User className="w-5 h-5" />
            <span className="hidden lg:inline">Login / Signup</span>
          </Link>

          <Link to="/cart" className="relative">
            <Button variant="ghost" size="icon" className="relative">
              <ShoppingCart className="w-5 h-5" />
              <span className="absolute -top-1 -right-1 bg-accent text-accent-foreground text-xs w-5 h-5 rounded-full flex items-center justify-center font-medium">
                {cartCount}
              </span>
            </Button>
          </Link>
        </div>
      </div>

      {/* Mobile Search */}
      <div className="md:hidden px-4 pb-4">
        <div className="relative">
          <Input
            type="search"
            placeholder="Search for products..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="pr-10 h-10 rounded-full"
          />
          <Button
            size="icon"
            variant="ghost"
            className="absolute right-1 top-1/2 -translate-y-1/2 h-8 w-8 rounded-full"
          >
            <Search className="w-4 h-4 text-muted-foreground" />
          </Button>
        </div>
      </div>
    </header>
  );
};

export default Header;
