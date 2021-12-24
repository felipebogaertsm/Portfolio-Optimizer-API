# Author: Felipe Bogaerts de Mattos
# email: felipebogaerts@gmail.com
# License: MIT

"""
Plotter file for the Markowitz optimizer.

Reads from the return_variance_data.json file.
"""

import plotly.graph_objects as go
import numpy as np
import pandas as pd

# Forming lists:
dataframe = pd.read_json("return_variance_data.json")
return_array = dataframe["Return"].to_numpy()
variance_array = dataframe["Variance"].to_numpy()

# Plotting:
fig = go.Figure(
    data=go.Scatter(
        y=return_array * 100,
        x=np.sqrt(variance_array) * 100,
        mode="markers",
        marker=dict(
            color=return_array / np.sqrt(variance_array),
            showscale=True,
            size=7,
            line=dict(width=1),
            colorscale="RdBu",
            colorbar=dict(title="Sharpe<br>Ratio"),
        ),
    )
)
fig.update_layout(
    xaxis=dict(title="Annualised Risk (Volatility)"),
    yaxis=dict(title="Annualised Return"),
    title="Sample of Random Portfolios",
)
fig.update_layout(coloraxis_colorbar=dict(title="Sharpe Ratio"))

fig.show()
