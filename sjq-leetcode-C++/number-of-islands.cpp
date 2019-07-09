class Solution {
public:
    int x, y;
    void dfs(int xx, int yy, vector<vector<char>>& grid){
        if(xx <0|| xx >=x || yy <0 || yy >=y || grid[xx][yy] =='0' )
            return ;
        grid[xx][yy] = '0';
        dfs(xx+1, yy, grid);
        dfs(xx, yy+1, grid);
        dfs(xx-1, yy, grid);
        dfs(xx, yy-1, grid);
    }
    
    int numIslands(vector<vector<char>>& grid) {
        if(grid.empty()) return 0;
        int res = 0;
        x = grid.size();
        y = grid[0].size();
        for(int i = 0; i < x; i++){
            for(int j = 0; j < y; j++){
                if(grid[i][j] == '1'){
                    dfs(i, j, grid);
                    res++;
                }
            }
        }
        return res;
    }
};
