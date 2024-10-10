export interface Category {
    ic_id: number;
    ic_name: string;
    ic_description: string;
    ic_picture_url: string;
    total_items_count: number;
};


export interface User {
    u_id: number;
    u_email: string;
    u_uid: string;
    u_register_date: string;
    u_is_disabled: boolean;
};



export interface Item {
    i_id: number,
    i_title: string,
    i_description: string,
    i_price: number,
    i_quantity: number,
    i_picture_url: string,
    i_file_path: string,
    i_is_disabled: boolean,
    i_release_date: string,
};

export interface SubCategory {
    isc_id: number;
    isc_name: string;
    isc_description: string;
    isc_picture_url: string;
};
