'use client';

import { useAuthState } from "react-firebase-hooks/auth";
import { auth } from "@/app/firebase/config";
import { useStoreModal } from "@/hooks/use-store-modal";
import { useEffect } from "react";



const SetupPage = () => {
    const [user] = useAuthState(auth);
    console.log({ user })

    const onOpen = useStoreModal((state) => state.onOpen);
    const isOpen = useStoreModal((state) => state.isOpen);

    useEffect(() => {
        if (!isOpen) {
            onOpen();
        }
    }, [onOpen, isOpen]);


    return null;
}


export default SetupPage;