


import { SubCategory, Category } from '@/app/schema/schema';
import { SubCategoryForm } from './components/sub-category-form';
import axios from 'axios';
import { cookies } from 'next/headers';

const SubCategoryPage = async ({ params }: { params: { subcategoryId: string } }) => {


    let subCategoryData: SubCategory | null = null;
    let categoryData: Category[] = [];
    let initialCategoryData: Category | null = null;

    const cookieStore = cookies();
    const token = cookieStore.get('token')?.value;


    if (params.subcategoryId && params.subcategoryId !== 'new') {
        try {
            const response = await axios.get(`${process.env.PB_PATH_API}/items/sub/category/${params.subcategoryId}`);

            if (response.status === 200) {
                subCategoryData = response.data;
            }
        } catch (error) {
        }

        try {

            const subCategory = await axios.get(`${process.env.PV_PATH_API}/items/category/link/${params.subcategoryId}?token=${token}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });


            if (subCategory.status === 200) {
                initialCategoryData = subCategory.data;
            }
        } catch (error) {
        }

    }

    try {
        const category = await axios.get(`${process.env.PB_PATH_API}/items/category`);
        if (category.status === 200) {
            categoryData = category.data;
        }
    } catch (error) {
        categoryData = [];
    }

    return (
        <div className="flex-col">
            <div className="flex-1 space-y-4 p-8 pt-6">
                <SubCategoryForm
                    initialCategory={initialCategoryData}
                    category={categoryData}
                    initialData={subCategoryData} />
            </div>
        </div>
    );
}

export default SubCategoryPage;
