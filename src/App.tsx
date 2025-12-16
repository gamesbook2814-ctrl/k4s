import { Toaster } from "@/components/ui/toaster";
import { Toaster as Sonner } from "@/components/ui/sonner";
import { TooltipProvider } from "@/components/ui/tooltip";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Index from "./pages/Index";
import Products from "./pages/Products";
import CategoryPage from "./pages/CategoryPage";
import ProductDetail from "./pages/ProductDetail";
import Cart from "./pages/Cart";
import Login from "./pages/Login";
import Help from "./pages/Help";
import About from "./pages/About";
import StoreLocator from "./pages/StoreLocator";
import SameDayDelivery from "./pages/SameDayDelivery";
import TrackOrder from "./pages/TrackOrder";
import NotFound from "./pages/NotFound";

const queryClient = new QueryClient();

const App = () => (
  <QueryClientProvider client={queryClient}>
    <TooltipProvider>
      <Toaster />
      <Sonner />
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Index />} />
          <Route path="/products" element={<Products />} />
          <Route path="/product/:id" element={<ProductDetail />} />
          <Route path="/cart" element={<Cart />} />
          <Route path="/login" element={<Login />} />
          <Route path="/help" element={<Help />} />
          <Route path="/about" element={<About />} />
          <Route path="/store-locator" element={<StoreLocator />} />
          <Route path="/same-day-delivery" element={<SameDayDelivery />} />
          <Route path="/track-order" element={<TrackOrder />} />
          
          {/* Category pages */}
          <Route path="/calendars-diaries" element={<CategoryPage />} />
          <Route path="/apparel" element={<CategoryPage />} />
          <Route path="/packaging" element={<CategoryPage />} />
          <Route path="/stationery" element={<CategoryPage />} />
          <Route path="/corporate-gifts" element={<CategoryPage />} />
          <Route path="/photo-gifts" element={<CategoryPage />} />
          <Route path="/drinkware" element={<CategoryPage />} />
          <Route path="/marketing-promo" element={<CategoryPage />} />
          <Route path="/gift-hampers" element={<CategoryPage />} />
          <Route path="/sample-kit" element={<CategoryPage />} />
          <Route path="/rewards-recognition" element={<CategoryPage />} />
          <Route path="/premium-products" element={<CategoryPage />} />
          
          <Route path="*" element={<NotFound />} />
        </Routes>
      </BrowserRouter>
    </TooltipProvider>
  </QueryClientProvider>
);

export default App;
