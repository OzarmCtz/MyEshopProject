import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { MainNav } from "@/components/main-nav";
import { Button } from "@/components/ui/button"
import { LogOut } from "lucide-react";
import { ThemeToggle } from "@/components/theme-toggle";
const Navbar = () => {
    return (
        <div className="border-b">
            <div className="flex h-16 items-center px-4">
                <div>
                    Admin-Panel
                </div>
                <MainNav className="mx-6" />
                <div className="ml-auto flex items-center space-x-4">
                    <Button variant="ghost" size="sm" className="flex items-center space-x-2">
                        <LogOut className="h-4 w-4" />
                        <span>Sign Out</span>
                    </Button>
                    <ThemeToggle />
                </div>
            </div>
        </div>
    )
    
}

export default Navbar;