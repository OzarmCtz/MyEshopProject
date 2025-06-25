import axios from "axios";

import { CategoryClient } from "./components/client";
import { CategoryColumn } from "./components/columns";

import { SubCategoryClient } from "../sub-categories/components/client";
import { SubCategoryColumn } from "../sub-categories/components/columns";

import { cookies } from 'next/headers';


const CategoriesPage = async () => {
    let formattedCategories: CategoryColumn[] = [];
    let formattedSubCategories: SubCategoryColumn[] = [];

    const cookieStore = cookies();
    const token = cookieStore.get('token')?.value;

    try {
        const categories = await axios.get(`${process.env.PV_PATH_API}/items/category?token=${token}`);

        formattedCategories = categories.data.map((item: any) => ({
            id: item.ic_id,
            name: item.ic_name,
            description: item.ic_description,
            pictureUrl: item.ic_picture_url,
            occurence: item.ic_on_isc,
            totalItemCount: item.total_items_count
        }));

    } catch (error) {
        formattedCategories = [];
    }

    try {
        const subCategories = await axios.get(`${process.env.PV_PATH_API}/items/sub/category?token=${token}`);

        formattedSubCategories = subCategories.data.map((item: any) => ({
            id: item.isc_id,
            name: item.isc_name,
            description: item.isc_description,
            pictureUrl: item.isc_picture_url,
            categoryName: item.ic_name,
            item_count: item.item_count
        }));

    } catch (error) {
        formattedSubCategories = [];
    }

    return (
        <div className="flex-col">
            <div className="flex-1 space-y-4 p-8 pt-6">
                <CategoryClient data={formattedCategories} />
            </div>
            <div className="flex-1 space-y-4 p-8 pt-6">
                <SubCategoryClient data={formattedSubCategories} />
            </div>
        </div>
    );
};

export default CategoriesPage;