import axios from "axios";

import { ItemClient } from "./components/client";
import { ItemColumn } from "./components/columns";
import { cookies } from 'next/headers';



const ItemsPage = async () => {
    let formattedItems: ItemColumn[] = [];
    const cookieStore = cookies();
    const token = cookieStore.get('token')?.value;


    try {
        const items = await axios.get(`${process.env.PV_PATH_API}/items?token=${token}`);

        formattedItems = items.data.map((item: any) => ({
            id: item.i_id,
            title: item.i_title,
            description: item.i_description,
            price: item.i_price,
            quantity: item.i_quantity,
            pictureUrl: item.i_picture_url,
            filePath: item.i_file_path,
            disabled: item.i_is_disabled,
            release_date: item.i_release_date,
            sub_category_name: item.sub_category_name,
        }));

    } catch (error) {
        formattedItems = [];
    }


    return (
        <div className="flex-col">
            <div className="flex-1 space-y-4 p-8 pt-6">
                <ItemClient data={formattedItems} />
            </div>
        </div>
    );
};

export default ItemsPage;