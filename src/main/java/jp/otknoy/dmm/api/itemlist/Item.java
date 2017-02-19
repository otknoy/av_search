package jp.otknoy.dmm.api.itemlist;

import lombok.Data;

@Data
public class Item {
    private String title;
    private String affiliateURL;
    private ImageUrl imageURL;
    private Prices prices;
    private String date;
    private Iteminfo iteminfo;
}
