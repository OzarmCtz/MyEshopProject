"use client"

import { Heading } from "@/components/ui/heading";
import { Button } from "@/components/ui/button";
import { Trash } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { z } from "zod";
import { useForm, FormProvider } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import toast from "react-hot-toast";
import { AlertModal } from "@/components/modals/alert-modal";
import { Router } from "next/router";
import { useRouter } from "next/navigation";
import { ApiAlert } from "@/components/ui/api-alert";
import { useOrigin } from "@/hooks/use-origin";

const formSchema = z.object({
    name: z.string().min(1),
});

type SettingsFormValues = z.infer<typeof formSchema>;

export const SettingsForm = () => {
    const [open, setOpen] = useState(false);
    const [loading, setLoading] = useState(false);
    const router = useRouter();
    const origin = useOrigin();


    const formMethods = useForm<SettingsFormValues>({
        resolver: zodResolver(formSchema),
    });

    const onSubmit = async (data: SettingsFormValues) => {

        try {
            console.log("test")
            setLoading(true);
            router.refresh()
            // do something
        } catch (error) {
            toast.error("Somethin went wrong");
        } finally {
            setLoading(false);
        }
        console.log(data);
    };

    const onDelete = async () => {
        try {
            setLoading(true);
            router.refresh()

            // do something
        } catch (error) {
            toast.error("Somethin went wrong");
        } finally {
            setLoading(false);
        }
    }

    return (
        <>
            <AlertModal
                isOpen={open}
                onClose={() => { setOpen(false) }}
                onConfirm={() => { }}
                loading={loading}
            />
            <div className="flex items-center justify-between">
                <Heading
                    title="Settings"
                    description="Manage store preferences"
                />
                <Button
                    disabled={loading}
                    variant="destructive"
                    size="icon"
                    onClick={() => { setOpen(true) }}
                >
                    <Trash className="h-4 w-4" />
                </Button>
            </div>
            <Separator />
            <FormProvider {...formMethods}>
                <form onSubmit={formMethods.handleSubmit(onSubmit)} className="space-y-8 w-full">
                    <div className="grid grid-cols-3 gap-8">
                        <FormField
                            control={formMethods.control}
                            name="name"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Name</FormLabel>
                                    <FormControl>
                                        <Input disabled={loading} placeholder="Store Name" {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                    </div>
                    <Button disabled={loading} className="ml-auto" type="submit">
                        Save Changes
                    </Button>
                </form>
            </FormProvider>
            <Separator />
            <ApiAlert
                title="PUBLIC_API_URL"
                description={`${process.env.NEXT_PUBLIC_PB_PATH_API}`}
                variant="public"
            />
            <ApiAlert
                title="PRIVATE_API_URL"
                description={`${process.env.NEXT_PUBLIC_PV_PATH_API}`}
                variant="admin"
            />
        </>
    );
};
