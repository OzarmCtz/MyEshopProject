"use client"
import { PopoverTrigger } from "@/components/ui/popover";
import { useStoreModal } from "@/hooks/use-store-modal";
import { useParams } from "next/navigation";
import { useRouter } from "next/router";

type PopoverTriggerPropos = React.ComponentPropsWithoutRef<typeof PopoverTrigger>



interface StoreSwitcherProps extends PopoverTriggerPropos {
    items: [

    ]
}

export default function StoreSwitcher({
    className,
    items = [],
}: StoreSwitcherProps) {
    const sotreModal = useStoreModal();
    const params = useParams();
    const router = useRouter();

    const formattedItems = items.map((item) => ({
        label: item,
        value: item,
    }));
    return (
        <div>
            Store Switcher
        </div>
    );
};